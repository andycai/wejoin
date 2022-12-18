package system

import (
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
)

type ActivitySystem struct {
}

var Activity = new(ActivitySystem)

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
