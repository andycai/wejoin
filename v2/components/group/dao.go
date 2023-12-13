package group

import (
	"errors"
	"time"

	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/library/math"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
	"github.com/andycai/axe-fiber/v2/model"
)

type GroupDao struct{}

var Dao = new(GroupDao)

var newErr = enum.GetError

// getGroupMemberLimit 暂时返回默认数量，以后会根据等级提升数量
func getGroupMemberLimit(gid int32) int32 {
	return enum.DefaultGroupMemmberCount
}

// isGroupManager 是否管理员（包括群主）
func isGroupManager(gid, uid int32) bool {
	gm := dao.GroupMember
	member, err := gm.Where(gm.GroupID.Eq(gid), gm.UserID.Eq(uid)).Take()

	return err == nil && member != nil && member.Position >= enum.PositionGroupManager
}

// isGroupOwner 是否群主
func isGroupOwner(gid, uid int32) bool {
	gm := dao.GroupMember
	member, err := gm.Where(gm.GroupID.Eq(gid), gm.UserID.Eq(uid)).Take()

	return err == nil && member != nil && member.Position == enum.PositionGroupOwner
}

// getGroup 返回群组
func getGroup(gid int32) *model.Group {
	g := dao.Group
	if group, err := g.Where(g.ID.Eq(gid)).Take(); err == nil {
		return group
	}

	return nil
}

// absent 群组不存在
func absent(gid int32) bool {
	return getGroup(gid) == nil
}

// end 私有方法

// GetInfo 返回群组信息
func (gd GroupDao) GetInfo(gid int32) (*comp.APIGroup, error) {
	g := dao.Group
	info := &comp.APIGroup{}
	err := g.Where(g.ID.Eq(gid)).Scan(info)
	if info.ID == 0 {
		err = newErr(enum.ErrorTextGroupNotFound)
	}

	return info, err
}

// GetGroupsByUserID 返回群组列表
func (gd GroupDao) GetGroupsByUserID(uid int32) ([]*comp.APIGroup, error) {
	ids := make([]int32, 0)
	m := dao.GroupMember
	err := m.Select(m.GroupID).Where(m.UserID.Eq(uid)).Scan(&ids)
	if err != nil {
		return nil, err
	}

	g := dao.Group
	groups := make([]*comp.APIGroup, len(ids))
	err = g.Where(g.ID.In(ids...)).Scan(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// GetGroups 返回最近的群组列表
func (gd GroupDao) GetGroups(page int, num int) ([]*comp.APIGroup, error) {
	g := dao.Group
	list := make([]*comp.APIGroup, 0)
	max := math.Max[int]
	page = max(page-1, 0)
	if num <= 0 {
		num = enum.DefaultGroupCount
	}
	err := g.Offset(page * num).Limit(num).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// GetApplyList 返回群组的申请入群列表
func (gd GroupDao) GetApplyList(gid int32) ([]comp.APIGroupApplication, error) {
	ga := dao.GroupApplication
	list := make([]comp.APIGroupApplication, 0)
	err := ga.Where(ga.GroupID.Eq(gid)).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Create 创建群组
func (gd GroupDao) Create(name, addr string) error {
	g := dao.Group

	// 检查名字是否存在
	count, err := g.Where(g.Name.Eq(name)).Count()
	if err == nil && count > 0 {
		return newErr(enum.ErrorTextGroupNameExists)
	}

	group := model.Group{}
	group.Name = name
	group.Addr = addr
	group.Level = 1
	group.Scores = 0
	group.Notice = "Welcome!"
	err = g.Create(&group)

	return err
}

// Apply 申请加入群组
func (gd GroupDao) Apply(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	g := dao.Group
	_, err := g.Where(g.ID.Eq(gid)).Take()
	if err != nil {
		return err
	}

	ga := dao.GroupApplication
	_, err = ga.Where(ga.GroupID.Eq(gid)).Where(ga.UserID.Eq(uid)).Take()

	if err != nil {
		groupApp := &model.GroupApplication{}
		groupApp.GroupID = gid
		groupApp.UserID = uid
		err = ga.Omit(ga.Deleted, ga.DeleteAt).Create(groupApp)

		return err
	}

	return errors.New("application exists")
}

// Approve 批准加入
func (gd GroupDao) Approve(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.GroupMember
	count, err := gm.Where(gm.GroupID.Eq(gid)).Count()
	if err != nil {
		return err
	}
	if count >= int64(getGroupMemberLimit(gid)) {
		return newErr(enum.ErrorTextGroupMemberFull)
	}

	ga := dao.GroupApplication
	// 管理员才能操作
	if isGroupManager(gid, mid) {
		rs, err := ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Take()
		if err == nil && rs != nil {
			if rs.Deleted == 0 {
				// 加入群组成员
				gm := dao.GroupMember
				member := &model.GroupMember{}
				member.GroupID = gid
				member.UserID = uid
				member.Position = enum.PositionGroupMember
				member.Scores = 0
				member.EnterAt = time.Now()
				member.Alias_ = ""
				err := gm.Create(member)
				if err != nil {
					return err
				}

				//  更新申请状态
				_, err = ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Update(ga.Deleted, 1)
				return err
			}
		}
	}

	return newErr(enum.ErrorTextGroupManagerOp)
}

// Refuse 拒绝
func (gd GroupDao) Refuse(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	ga := dao.GroupApplication
	// 管理员才能操作
	if isGroupManager(gid, mid) {
		rs, err := ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Take()
		if err == nil && rs != nil {
			//  更新申请状态
			_, err = ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Update(ga.Deleted, 1)
		}
		return err
	}

	return newErr(enum.ErrorTextGroupManagerOp)
}

// Promote 提升管理员
func (gd GroupDao) Promote(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.GroupMember
	// 群管理员数量已满
	count, err := gm.Where(gm.GroupID.Eq(gid), gm.Position.Eq(enum.PositionGroupManager)).Count()
	if err != nil {
		return err
	}
	if count >= enum.DefaultGroupManagerCount {
		return newErr(enum.ErrorTextGroupManagerFull)
	}

	// 群组可以提升管理员
	if isGroupOwner(gid, mid) {
		result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Update(gm.Position, enum.PositionGroupManager)
		if err != nil {
			return err
		}
		return result.Error
	}

	return newErr(enum.ErrorTextGroupPromote)
}

// Transfer 转让群主
func (gd GroupDao) Transfer(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	if isGroupOwner(gid, mid) {
		gm := dao.GroupMember
		result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Update(gm.Position, enum.PositionGroupOwner)
		if err == nil {
			result, err := gm.Where(gm.UserID.Eq(mid), gm.GroupID.Eq(gid)).Update(gm.Position, enum.PositionGroupMember)
			if err != nil {
				return err
			}
			return result.Error
		}
		return result.Error
	}

	return newErr(enum.ErrorTextGroupTransfer)
}

// Fire 踢出群组
func (gd GroupDao) Fire(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 管理员才能踢人
	if isGroupManager(gid, mid) {
		gm := dao.GroupMember
		result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Delete()
		if err == nil {
			return err
		}

		return result.Error
	}

	return newErr(enum.ErrorTextGroupFireMember)
}

// Remove 删除群组
func (gd GroupDao) Remove(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 群主可以删除群组
	if isGroupOwner(gid, uid) {
		g := dao.Group
		result, err := g.Where(g.ID.Eq(gid)).Delete()
		if err != nil {
			return err
		}
		return result.Error
	}

	return newErr(enum.ErrorTextGroupRemove)
}

// Quit 退出群组
func (gd GroupDao) Quit(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.GroupMember
	result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Delete()
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateName 更新群名字
func (gd GroupDao) UpdateName(gid, uid int32, name string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isGroupManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	// TODO: 改名次数限制

	// 是否存在名称
	count, err := g.Where(g.Name.Eq(name), g.ID.Neq(gid)).Count()
	if err == nil && count > 0 {
		return newErr(enum.ErrorTextGroupNameExists)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Name, name)
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateLogo 更新 Logo
func (gd GroupDao) UpdateLogo(gid, uid int32, logo string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isGroupManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Logo, logo)
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateNotice 更新公告
func (gd GroupDao) UpdateNotice(gid, uid int32, notice string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isGroupManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Notice, notice)
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateAddr 更新地址
func (gd GroupDao) UpdateAddr(gid, uid int32, addr string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isGroupManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Addr, addr)
	if err != nil {
		return err
	}

	return result.Error
}
