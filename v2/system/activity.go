package system

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
)

type ActivitySystem struct {
}

var Activity = new(ActivitySystem)

// GetInfo 返回活动信息
func (as ActivitySystem) GetInfo(aid int32) (*comp.APIActivity, error) {
	a := dao.Q.Activity
	info := &comp.APIActivity{}
	err := a.Where(a.ID.Eq(aid)).Scan(info)
	if info.ID == 0 {
		err = newErr(enum.ErrorTextActivityNotFound)
	}

	return info, err
}

// GetActivitiesByUserID 返回活动列表
func (us ActivitySystem) GetActivitiesByUserID(uid int32) ([]*comp.APIActivity, error) {
	ids := make([]int32, 0)
	u := dao.Q.ActivityUser
	err := u.Select(u.ActivityID).Where(u.UserID.Eq(uid)).Scan(&ids)
	if err != nil {
		return nil, err
	}

	a := dao.Q.Activity
	activities := make([]*comp.APIActivity, len(ids))
	err = a.Where(a.ID.In(ids...)).Scan(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

// GetActivitiesByGroupID 返回群组创建的活动列表
func (as ActivitySystem) GetActivitiesByGroupID(gid int32) ([]*comp.APIActivity, error) {
	a := dao.Q.Activity
	list := make([]*comp.APIActivity, 0)
	err := a.Where(a.GroupID.Eq(gid)).Scan(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
