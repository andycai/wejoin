package activity

import (
	"github.com/andycai/wejoin/api/body"
	"github.com/andycai/wejoin/api/comp"
	"github.com/andycai/wejoin/api/dao"
	"github.com/andycai/wejoin/enum"
	"github.com/andycai/wejoin/library/math"
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

// GetByUserID get the activities by user id
func (as ActivityDao) GetByUserID(uid uint) ([]*model.Activity, error) {
	activitsies := make([]*model.Activity, 0)
	db.Raw(SqlQueryActivitiesByUserID, uid).Scan(&activitsies))

	return activitsies, nil
}

// GetActivitiesByGroupID 返回群组创建的活动列表
func (as ActivityDao) GetActivitiesByGroupID(gid uint) ([]*comp.APIActivity, error) {
	a := dao.Activity
	list := make([]*comp.APIActivity, 0)
	err := a.Where(a.GroupID.Eq(gid)).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// GetActivities 返回活动列表
func (as ActivityDao) GetActivities(page int, num int) ([]*comp.APIActivity, error) {
	a := dao.Activity
	list := make([]*comp.APIActivity, 0)
	max := math.Max[int]
	page = max(page-1, 0)
	if num <= 0 {
		num = enum.DefaultActivityCount
	}
	err := a.Offset(page * num).Limit(num).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Update 更新活动信息
func (as ActivityDao) Update(activity *body.Activity) error {
	a := dao.Activity
	_, err := a.Where(a.ID.Eq(activity.ID)).Updates(activity.ToModel())

	return err
}
