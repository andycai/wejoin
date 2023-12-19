package activity

import (
	"github.com/andycai/axe-fiber/api/body"
	"github.com/andycai/axe-fiber/api/comp"
	"github.com/andycai/axe-fiber/api/dao"
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/library/math"
)

type ActivityDao struct {
}

var Dao = new(ActivityDao)

var newErr = enum.GetError

// GetInfo 返回活动信息
func (as ActivityDao) GetInfo(aid uint) (*comp.APIActivity, error) {
	a := dao.Activity
	info := &comp.APIActivity{}
	err := a.Where(a.ID.Eq(aid)).Scan(info)
	if info.ID == 0 {
		err = newErr(enum.ErrorTextActivityNotFound)
	}

	return info, err
}

// GetActivitiesByUserID 返回活动列表
func (as ActivityDao) GetActivitiesByUserID(uid uint) ([]*comp.APIActivity, error) {
	ids := make([]uint, 0)
	u := dao.ActivityUser
	// TODO: 使用关联表方式
	err := u.Select(u.ActivityID).Where(u.UserID.Eq(uid)).Scan(&ids)
	if err != nil {
		return nil, err
	}

	a := dao.Activity
	activities := make([]*comp.APIActivity, len(ids))
	err = a.Where(a.ID.In(ids...)).Scan(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
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
