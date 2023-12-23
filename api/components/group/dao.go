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

// getMemberLimit return the maximum number of members
func getMemberLimit(gid uint) uint {
	return enum.DefaultGroupMemmberCount
}

// getManagerLimit return the maximum number of managers
func getManagerLimit(gid uint) uint {
	return enum.DefaultGroupManagerCount
}

// isManager is the manager of the group or not(including the owner)
func isManager(gid, uid uint) (err error) {
	member := &model.GroupMember{}
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndUserID, uid, gid).Take(member).Error

	if err != nil {
		return
	}

	if member == nil || member.Position >= enum.PositionGroupOwner {
		err = newErr(enum.ErrorTextGroupManagerOp)
	}

	return
}

// isOwner is the owner of the group or not
func isOwner(gid, uid uint) (err error) {
	member := &model.GroupMember{}
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndUserID, uid, gid).Take(member).Error

	if err != nil {
		return
	}

	if member == nil || member.Position != enum.PositionGroupOwner {
		err = newErr(enum.ErrorTextGroupManagerOp)
	}

	return
}

// existsGroup exists the group or not
func existsGroup(gid uint) (err error) {
	err = db.Raw(SqlQueryGroupByID, gid).First(&model.Group{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return nil
}

// existsMember exists the member of the group or not
func existsMember(gid, uid uint) (err error) {
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndUserID, gid, uid).First(&model.GroupMember{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return nil
}

// existsApplication exists application or not
func existsApplication(gid, uid uint) (err error) {
	ga := model.GroupApplication{}
	err = db.Raw(SqlQueryGroupApplicationsByGroupIDAndUserID, gid, uid).First(&ga).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	return
}

// canJoin can join the group or not
func canJoin(gid uint) (err error) {
	var count int64
	err = db.Raw(SqlQueryGroupMemberByGroupID, gid).Count(&count).Error
	if err != nil {
		return
	}

	if count >= int64(getMemberLimit(gid)) {
		err = newErr(enum.ErrorTextGroupMemberFull)
	}
	return
}

// canPromote can promote the member or not
func canPromote(gid uint) (err error) {
	var count int64
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndPosition, gid, enum.PositionGroupMember).Count(&count).Error
	if err != nil {
		return
	}

	if count >= int64(getManagerLimit(gid)) {
		err = newErr(enum.ErrorTextGroupManagerFull)
	}
	return
}

//#endregion

// GetByID get the groups by id
func (gd GroupDao) GetByID(gid uint) (*model.Group, error) {
	groupVo := model.Group{}
	db.Raw(SqlQueryGroupByID, gid).Scan(&groupVo)

	return &groupVo, nil
}

// GetByUserID get the group by user id
func (gd GroupDao) GetByUserID(uid uint) ([]*model.Group, error) {
	groups := make([]*model.Group, 0)
	// db.Raw(SqlGroupByUserID, uid).Find(&groups)
	db.Raw(SqlQueryGroupByUserID, uid).Scan(&groups)

	return groups, nil
}

// GetByPage get the groups by page
func (gd GroupDao) GetByPage(page int, pageSize int) ([]*model.Group, error) {
	groups := make([]*model.Group, 0)
	db.Raw(SqlQueryGroupByPage, pageSize, pageSize*(page-1)).Scan(&groups)

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
	err := db.Unscoped().Raw(SqlQueryGroupByID, gid).First(&model.Group{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// ExistsName group name exists or not
func (gd GroupDao) ExistsName(name string) error {
	err := db.Unscoped().Raw(SqlQueryGroupByName, name).First(&model.Group{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// GetApplictions return the applications of the group
func (gd GroupDao) GetApplictions(gid uint) ([]model.GroupApplication, error) {
	applications := make([]model.GroupApplication, 0)
	err := db.Raw(SqlQueryGroupApplicationsByGroupID).Scan(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

// Apply apply for the group
func (gd GroupDao) Apply(group *model.GroupApplication) error {
	err := existsGroup(group.GroupID)
	if err != nil {
		return err
	}

	err = canJoin(group.GroupID)
	if err != nil {
		return err
	}

	err = existsApplication(group.GroupID, group.UserID)
	if err == nil {
		return err
	}

	err = db.Create(&group).Error

	return err
}

// Approve approve the application
func (gd GroupDao) Approve(gid, uid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = canJoin(gid)
	if err != nil {
		return err
	}

	err = existsApplication(gid, mid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	// 事务：创建成员和删除申请
	tx := db.Begin()

	// 加入群组成员
	member := model.GroupMember{}
	member.GroupID = gid
	member.UserID = mid
	member.Position = enum.PositionGroupMember
	member.Scores = 0
	member.EnterAt = time.Now()
	member.Nick = ""
	err = tx.Create(&member).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Raw(SqlDeleteGroupApplicationByGroupIDAndUserID, gid, mid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

// Refuse refuse the application of the user
func (gd GroupDao) Refuse(gid, uid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsApplication(gid, mid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	err = db.Raw(SqlDeleteGroupApplicationByGroupIDAndUserID, gid, mid).Error

	return err
}

// Promote promote the member to the manager
func (gd GroupDao) Promote(gid, uid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsMember(gid, mid)
	if err != nil {
		return err
	}

	err = canPromote(gid)
	if err != nil {
		return err
	}

	err = isManager(gid, mid)
	if err == nil {
		return newErr(enum.ErrorTextGroupAlreadyManager)
	}

	err = isOwner(gid, uid)
	if err != nil {
		return err
	}

	err = db.Raw(SqlUpdateGroupMemberPositionByGroupIDAndUserID, enum.PositionGroupManager, gid, mid).Error

	return err
}

// Transfer transfer the owner to the other member
func (gd GroupDao) Transfer(gid, uid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsMember(gid, mid)
	if err != nil {
		return err
	}

	err = isOwner(gid, uid)
	if err != nil {
		return err
	}

	// transaction
	tx := db.Begin()

	err = tx.Raw(SqlUpdateGroupMemberPositionByGroupIDAndUserID, enum.PositionGroupMember, gid, uid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Raw(SqlUpdateGroupMemberPositionByGroupIDAndUserID, enum.PositionGroupOwner, gid, mid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// Fire kick the member out of the group
func (gd GroupDao) Fire(gid, uid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsMember(gid, mid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	err = db.Raw(SqlDeleteGroupMemberByGroupIDAndUserID, gid, mid).Error

	return err
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
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isOwner(gid, uid)
	if err == nil {
		return newErr(enum.ErrorTextGroupOwnerCannotQuit)
	}

	err = db.Raw(SqlDeleteGroupMemberByGroupIDAndUserID, gid, mid).Error

	return err
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
