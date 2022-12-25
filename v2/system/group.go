package system

import (
	"errors"
	"time"

	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/library/math"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
	"github.com/andycai/axe-fiber/v2/model"
)

type GroupSystem struct {
}

var Group = new(GroupSystem)

// begin 私有方法

// getGroupMemberLimit 暂时返回默认数量，以后会根据等级提升数量
func getGroupMemberLimit(gid int32) int32 {
	return enum.DefaultGroupMemmberCount
}

// isGroupManager 是否管理员（包括群主）
func isGroupManager(gid, uid int32) bool {
	gm := dao.Q.GroupMember
	member, err := gm.Where(gm.GroupID.Eq(gid), gm.UserID.Eq(uid)).Take()

	return err == nil && member != nil && member.Position >= enum.PositionGroupManager
}

// isGroupOwner 是否群主
func isGroupOwner(gid, uid int32) bool {
	gm := dao.Q.GroupMember
	member, err := gm.Where(gm.GroupID.Eq(gid), gm.UserID.Eq(uid)).Take()

	return err == nil && member != nil && member.Position == enum.PositionGroupOwner
}

// getGroup 返回群组
func getGroup(gid int32) *model.Group {
	g := dao.Q.Group
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
func (us GroupSystem) GetInfo(gid int32) (*comp.APIGroup, error) {
	g := dao.Q.Group
	info := &comp.APIGroup{}
	err := g.Where(g.ID.Eq(gid)).Scan(info)
	if info.ID == 0 {
		err = newErr(enum.ErrorTextGroupNotFound)
	}

	return info, err
}

// GetGroupsByUserID 返回群组列表
func (us GroupSystem) GetGroupsByUserID(uid int32) ([]*comp.APIGroup, error) {
	ids := make([]int32, 0)
	m := dao.Q.GroupMember
	err := m.Select(m.GroupID).Where(m.UserID.Eq(uid)).Scan(&ids)
	if err != nil {
		return nil, err
	}

	g := dao.Q.Group
	groups := make([]*comp.APIGroup, len(ids))
	err = g.Where(g.ID.In(ids...)).Scan(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// GetGroups 返回最近的群组列表
func (us GroupSystem) GetGroups(page int, num int) ([]*comp.APIGroup, error) {
	g := dao.Q.Group
	groups := make([]*comp.APIGroup, 0)
	max := math.Max[int]
	page = max(page-1, 0)
	if num <= 0 {
		num = enum.DefaultGroupCount
	}
	err := g.Offset(page * num).Limit(num).Scan(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

// GetApplyList 返回群组的申请入群列表
func (us GroupSystem) GetApplyList(gid int32) ([]comp.APIGroupApplication, error) {
	ga := dao.Q.GroupApplication
	list := make([]comp.APIGroupApplication, 0)
	err := ga.Where(ga.GroupID.Eq(gid)).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Create 创建群组
func (us GroupSystem) Create() error {
	//
	return nil
}

// Apply 申请加入群组
func (us GroupSystem) Apply(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	g := dao.Q.Group
	_, err := g.Where(g.ID.Eq(gid)).Take()
	if err != nil {
		return err
	}

	ga := dao.Q.GroupApplication
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
func (us GroupSystem) Approve(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	ga := dao.Q.GroupApplication
	rs, err := ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Take()
	if err == nil && rs != nil {
		if rs.Deleted == 0 {
			// 加入群组成员
			gm := dao.Q.GroupMember
			member := &model.GroupMember{}
			member.GroupID = gid
			member.UserID = uid
			member.Position = enum.PositionGroupMember
			member.Scores = 0
			member.EnterAt = time.Now()
			member.Alias_ = ""
			member.Avatar = ""
			err := gm.Create(member)
			if err != nil {
				return err
			}

			//  更新申请状态
			_, err = ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Update(ga.Deleted, 1)
			return err
		}
	}

	return errors.New("approve failed")
}

// Refuse 拒绝
func (us GroupSystem) Refuse(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	ga := dao.Q.GroupApplication
	rs, err := ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Take()
	if err == nil && rs != nil {
		//  更新申请状态
		_, err = ga.Where(ga.GroupID.Eq(gid), ga.UserID.Eq(uid)).Update(ga.Deleted, 1)
	}

	return err
}

// Promote 提升管理员
func (us GroupSystem) Promote(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.Q.GroupMember
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

// Transfer 转让群组
func (us GroupSystem) Transfer(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	if isGroupOwner(gid, mid) {
		gm := dao.Q.GroupMember
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
func (us GroupSystem) Fire(gid, uid, mid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 管理员才能踢人
	if isGroupManager(gid, mid) {
		gm := dao.Q.GroupMember
		result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Delete()
		if err == nil {
			return err
		}

		return result.Error
	}

	return newErr(enum.ErrorTextGroupFireMember)
}

// Remove 删除群组
func (us GroupSystem) Remove(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 群主可以删除群组
	if isGroupOwner(gid, uid) {
		g := dao.Q.Group
		result, err := g.Where(g.ID.Eq(gid)).Delete()
		if err != nil {
			return err
		}
		return result.Error
	}

	return newErr(enum.ErrorTextGroupRemove)
}

// Quit 退出群组
func (us GroupSystem) Quit(gid, uid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.Q.GroupMember
	result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Delete()
	if err != nil {
		return err
	}

	return result.Error
}

// Update 更新群组资料
func (us GroupSystem) Update(gid int32) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	return nil
}
