package activity

import (
	"github.com/andycai/wejoin/api/body"
	"github.com/andycai/wejoin/api/dao"
	"github.com/andycai/wejoin/enum"
	"github.com/andycai/wejoin/model"
)

type ActivityDao struct {
}

var Dao = new(ActivityDao)

var newErr = enum.GetError

// GetByID get the activity by the activity id
func (as ActivityDao) GetByID(aid uint) (*model.Activity, error) {
	activityVo := model.Activity{}
	db.Raw(SqlQueryActivityByID, aid).Scan(&activityVo)

	return &activityVo, nil
}

// GetByOrganizerUserID get the activities by organizer user id
func (as ActivityDao) GetByOrganizerUserID(uid uint) ([]*model.Activity, error) {
	activitsies := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByOrganizerUserID, uid).Scan(&activitsies)

	return activitsies, nil
}

// GetByUserID get the activities by user id
func (as ActivityDao) GetByUserID(uid uint) ([]*model.Activity, error) {
	activitsies := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByUserID, uid).Scan(&activitsies)

	return activitsies, nil
}

// GetByGroupID get the activities by group id
func (as ActivityDao) GetByGroupID(gid uint) ([]*model.Activity, error) {
	activitsies := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByGroupID, gid).Scan(&activitsies)

	return activitsies, nil
}

// GetByPage get the activities by page
func (as ActivityDao) GetByPage(page int, pageSize int) ([]*model.Activity, error) {
	activities := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByPage, page, pageSize*(page-1)).Scan(&activities)

	return activities, nil
}

// Update 更新活动信息
func (as ActivityDao) Update(activity *body.Activity) error {
	a := dao.Activity
	_, err := a.Where(a.ID.Eq(activity.ID)).Updates(activity.ToModel())

	return err
}
