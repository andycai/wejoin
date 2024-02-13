package group

import (
	"errors"
	"time"

	"github.com/andycai/wejoin/api/components/activity"
	"github.com/andycai/wejoin/constant"
	"github.com/andycai/wejoin/model"
	"gorm.io/gorm"
)

type GroupDao struct{}

var Dao = new(GroupDao)

var newErr = constant.GetError

//#region private methods

// getMemberLimit return the maximum number of members
func getMemberLimit(gid uint) uint {
	return constant.DefaultGroupMemmberCount
}

// getManagerLimit return the maximum number of managers
func getManagerLimit(gid uint) uint {
	return constant.DefaultGroupManagerCount
}

// isManager is the manager of the group or not(including the owner)
func isManager(gid, uid uint) (err error) {
	member := &model.GroupMember{}
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndUserID, uid, gid).Take(member).Error

	if err != nil {
		return
	}

	if member == nil || member.Position >= constant.PositionGroupOwner {
		err = newErr(constant.ErrorTextGroupManagerOp)
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

	if member == nil || member.Position != constant.PositionGroupOwner {
		err = newErr(constant.ErrorTextGroupManagerOp)
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
	err = db.Raw(SqlQueryGroupApplicationByGroupIDAndUserID, gid, uid).First(&ga).Error
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
		err = newErr(constant.ErrorTextGroupMemberFull)
	}
	return
}

// canPromote can promote the member or not
func canPromote(gid uint) (err error) {
	var count int64
	err = db.Raw(SqlQueryGroupMemberByGroupIDAndPosition, gid, constant.PositionGroupMember).Count(&count).Error
	if err != nil {
		return
	}

	if count >= int64(getManagerLimit(gid)) {
		err = newErr(constant.ErrorTextGroupManagerFull)
	}
	return
}

//#endregion

// GetByID get the group by the group id
func (gd GroupDao) GetByID(gid uint) (*model.Group, error) {
	groupVo := model.Group{}
	db.Raw(SqlQueryGroupByID, gid).Scan(&groupVo)

	return &groupVo, nil
}

// GetByUserID get the groups by user id
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
		return newErr(constant.ErrorTextGroupNameExists)
	}

	group.Level = 1
	group.Scores = 0
	group.Notice = ""
	err := db.Create(group).Error

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

// GetApplictionsByGroupID return the applications of the group
func (gd GroupDao) GetApplictionsByGroupID(gid uint) ([]model.GroupApplication, error) {
	applications := make([]model.GroupApplication, 0)
	err := db.Raw(SqlQueryGroupApplicationsByGroupID).Scan(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

// Apply apply for the group
func (gd GroupDao) Apply(gid, uid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = canJoin(gid)
	if err != nil {
		return err
	}

	err = existsApplication(gid, uid)
	if err == nil {
		return err
	}

	groupApplication := model.GroupApplication{}
	groupApplication.GroupID = gid
	groupApplication.UserID = uid
	err = db.Create(&groupApplication).Error

	return err
}

// Cancel cancel the application of the group
func (gd GroupDao) Cancel(gid, uid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsApplication(gid, uid)
	if err != nil {
		return err
	}

	err = db.Exec(SqlDeleteGroupApplicationByGroupIDAndUserID, gid, uid).Error

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
	member.Position = constant.PositionGroupMember
	member.Scores = 0
	member.EnterAt = time.Now()
	member.Nick = ""
	err = tx.Create(&member).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Exec(SqlDeleteGroupApplicationByGroupIDAndUserID, gid, mid).Error
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

	err = db.Exec(SqlDeleteGroupApplicationByGroupIDAndUserID, gid, mid).Error

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
		return newErr(constant.ErrorTextGroupAlreadyManager)
	}

	err = isOwner(gid, uid)
	if err != nil {
		return err
	}

	err = db.Exec(SqlUpdateGroupMemberPositionByGroupIDAndUserID, constant.PositionGroupManager, gid, mid).Error

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

	err = tx.Exec(SqlUpdateGroupMemberPositionByGroupIDAndUserID, constant.PositionGroupMember, gid, uid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Exec(SqlUpdateGroupMemberPositionByGroupIDAndUserID, constant.PositionGroupOwner, gid, mid).Error
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

	err = db.Exec(SqlDeleteGroupMemberByGroupIDAndUserID, gid, mid).Error

	return err
}

// Remove remove the group
func (gd GroupDao) Remove(gid, uid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isOwner(gid, uid)
	if err != nil {
		return err
	}

	// transaction

	tx := db.Begin()

	// delete the applications
	err = db.Exec(SqlDeleteGroupApplicationByGroupID, gid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// delete the members
	err = tx.Exec(SqlDeleteGroupMemberByGroupID, gid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// TODO delete the activities
	err = tx.Exec(activity.SqlDeleteActivitiesByGroupID, gid).Error

	// delete the group
	err = tx.Exec(SqlDeleteGroupnByID, gid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// Quit quit the group
func (gd GroupDao) Quit(gid, mid uint) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = existsMember(gid, mid)
	if err != nil {
		return err
	}

	err = isOwner(gid, mid)
	if err == nil {
		return newErr(constant.ErrorTextGroupOwnerCannotQuit)
	}

	err = db.Exec(SqlDeleteGroupMemberByGroupIDAndUserID, gid, mid).Error

	return err
}

// UpdateName update the name of the group
func (gd GroupDao) UpdateName(gid, uid uint, name string) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	// TODO: 改名次数限制

	if gd.ExistsName(name) != nil {
		return newErr(constant.ErrorTextGroupNameExists)
	}

	err = db.Exec(SqlUpdateGroupNameByID, name, gid).Error

	return err
}

// UpdateLogo update the logo of the group
func (gd GroupDao) UpdateLogo(gid, uid uint, logo string) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	err = db.Exec(SqlUpdateGroupLogoByID, logo, gid).Error

	return err
}

// UpdateNotice update the logo of the group
func (gd GroupDao) UpdateNotice(gid, uid uint, notice string) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	err = db.Exec(SqlUpdateGroupLogoByID, notice, gid).Error

	return err
}

// UpdateAddress update the address of the group
func (gd GroupDao) UpdateAddress(gid, uid uint, addr string) error {
	err := existsGroup(gid)
	if err != nil {
		return err
	}

	err = isManager(gid, uid)
	if err != nil {
		return err
	}

	err = db.Exec(SqlUpdateGroupAddrByID, addr, gid).Error

	return err
}
