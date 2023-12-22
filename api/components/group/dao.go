package group

import (
	"errors"
	"time"

	"github.com/andycai/axe-fiber/api/dao"
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/model"
	"gorm.io/gorm"
)

type GroupDao struct{}

var Dao = new(GroupDao)

var newErr = enum.GetError

//#region private methods

// getMemberLimit 暂时返回默认数量，以后会根据等级提升数量
func getMemberLimit(gid uint) uint {
	return enum.DefaultGroupMemmberCount
}

// isManager 是否管理员（包括群主）
func isManager(gid, uid uint) bool {
	member := &model.GroupMember{}
	err := db.Raw(SqlGroupMemberByGroupIDAndUserID, uid, gid).Take(member).Error

	return err == nil && member != nil && member.Position >= enum.PositionGroupOwner
}

// isOwner 是否群主
func isOwner(gid, uid uint) bool {
	member := &model.GroupMember{}
	err := db.Raw(SqlGroupMemberByGroupIDAndUserID, uid, gid).Take(member).Error

	return err == nil && member != nil && member.Position == enum.PositionGroupOwner
}

// absent group doesn't exist
func absent(gid uint) bool {
	return Dao.Exists(gid) != nil
}

//#endregion

// GetByID get the groups by id
func (gd GroupDao) GetByID(gid uint) (*model.Group, error) {
	groupVo := model.Group{}
	db.Raw(SqlGroupByID, gid).Scan(&groupVo)

	return &groupVo, nil
}

// GetByUserID get the group by user id
func (gd GroupDao) GetByUserID(uid uint) ([]*model.Group, error) {
	groups := make([]*model.Group, 0)
	// db.Raw(SqlGroupByUserID, uid).Find(&groups)
	db.Raw(SqlGroupByUserID, uid).Scan(&groups)

	return groups, nil
}

// GetByPage get the groups by page
func (gd GroupDao) GetByPage(page int, pageSize int) ([]*model.Group, error) {
	groups := make([]*model.Group, 0)
	db.Raw(SqlGroupByPage, pageSize, pageSize*(page-1)).Scan(&groups)

	return groups, nil
}

// Create create a new group
func (gd GroupDao) Create(group *model.Group) error {
	if gd.ExistsName(group.Name) != nil {
		return newErr(enum.ErrorTextGroupNameExists)
	}

	group.Level = 1
	group.Scores = 0
	group.Notice = ""
	err := db.Create(group).Error
	if err != nil {
		return err
	}

	return err
}

// Exists group exists or not
func (gd GroupDao) Exists(gid uint) error {
	err := db.Unscoped().Raw(SqlGroupByID, gid).First(&model.Group{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// ExistsName group name exists or not
func (gd GroupDao) ExistsName(name string) error {
	err := db.Unscoped().Raw(SqlGroupByName, name).First(&model.Group{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// GetApplictions return the applications of the group
func (gd GroupDao) GetApplictions(gid uint) ([]model.GroupApplication, error) {
	applications := make([]model.GroupApplication, 0)
	err := db.Raw(SqlGroupApplicationsByGroupID).Scan(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

// Apply apply for the group
func (gd GroupDao) Apply(group *model.GroupApplication) error {
	if absent(group.GroupID) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	groupVo := model.Group{}

	err := db.Unscoped().Raw(SqlGroupApplicationsByGroupIDAndUserID, group.GroupID, group.UserID).First(&groupVo).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = db.Create(&group).Error

	return err
}

// Approve 批准加入
func (gd GroupDao) Approve(gid, uid, mid uint) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	gm := dao.GroupMember
	count, err := gm.Where(gm.GroupID.Eq(gid)).Count()
	if err != nil {
		return err
	}
	if count >= int64(getMemberLimit(gid)) {
		return newErr(enum.ErrorTextGroupMemberFull)
	}

	ga := dao.GroupApplication
	// 管理员才能操作
	if isManager(gid, mid) {
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
func (gd GroupDao) Refuse(gid, uid, mid uint) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	ga := dao.GroupApplication
	// 管理员才能操作
	if isManager(gid, mid) {
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
func (gd GroupDao) Promote(gid, uid, mid uint) error {
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
	if isOwner(gid, mid) {
		result, err := gm.Where(gm.UserID.Eq(uid), gm.GroupID.Eq(gid)).Update(gm.Position, enum.PositionGroupManager)
		if err != nil {
			return err
		}
		return result.Error
	}

	return newErr(enum.ErrorTextGroupPromote)
}

// Transfer 转让群主
func (gd GroupDao) Transfer(gid, uid, mid uint) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	if isOwner(gid, mid) {
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
func (gd GroupDao) Fire(gid, uid, mid uint) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 管理员才能踢人
	if isManager(gid, mid) {
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
func (gd GroupDao) Remove(gid, uid uint) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}

	// 群主可以删除群组
	if isOwner(gid, uid) {
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
func (gd GroupDao) Quit(gid, uid uint) error {
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
func (gd GroupDao) UpdateName(gid, uid uint, name string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isManager(gid, uid) {
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
func (gd GroupDao) UpdateLogo(gid, uid uint, logo string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Logo, logo)
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateNotice 更新公告
func (gd GroupDao) UpdateNotice(gid, uid uint, notice string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Notice, notice)
	if err != nil {
		return err
	}

	return result.Error
}

// UpdateAddr 更新地址
func (gd GroupDao) UpdateAddr(gid, uid uint, addr string) error {
	if absent(gid) {
		return newErr(enum.ErrorTextGroupNotFound)
	}
	g := dao.Group

	if isManager(gid, uid) {
		return newErr(enum.ErrorTextGroupManagerOp)
	}

	result, err := g.Where(g.ID.Eq(gid)).Update(g.Addr, addr)
	if err != nil {
		return err
	}

	return result.Error
}
