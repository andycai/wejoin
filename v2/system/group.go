package system

import (
	"errors"

	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/library/math"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
)

type GroupSystem struct {
}

var Group = new(GroupSystem)

// GetInfo 返回群组信息
func (us GroupSystem) GetInfo(gid int32) (*comp.APIGroup, error) {
	g := dao.Q.Group
	info := &comp.APIGroup{}
	err := g.Where(g.ID.Eq(gid)).Scan(info)
	if info.ID == 0 {
		err = errors.New("not found group data")
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
	err := g.Debug().Offset(page * num).Limit(num).Scan(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (us GroupSystem) Exists(id int32) bool {
	return true
}
