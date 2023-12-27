package activity

import (
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
func (as ActivityDao) GetByPage(page uint, pageSize uint) ([]*model.Activity, error) {
	activities := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByPage, page, pageSize*(page-1)).Scan(&activities)

	return activities, nil
}

// Create create a new activity
func (as ActivityDao) Create(activity *model.Activity) error {
	err := db.Create(activity).Error

	return err
}

// Update update the activity details
func (as ActivityDao) Update(activity *model.Activity) error {
	err := db.Model(activity).Select("title", "description", "ahead", "begin_at", "end_at").Updates(activity).Error

	return err
}
