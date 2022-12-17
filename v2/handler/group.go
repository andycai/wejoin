package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/dao"
)

type GroupHandler struct{}

var Group = new(GroupHandler)

// GetGroupsById 获取单个群组信息
func (g GroupHandler) GetGroupById(c *Ctx) error {
	gid := I32(c, "gid")
	q := dao.Q.Group
	group, err := q.Where(q.ID.Eq(gid)).Take()
	if err != nil {
		return Err(c, enum.ErrGroupNotFound)
	}

	return Ok(c, group)
}

// GetGroupsByUserId 获取当前登录用户的群组列表
func (g GroupHandler) GetGroupsByUserId(c *Ctx) error {
	uid := I32(c, "uid")
	// q := dao.Q.Group
	m := dao.Q.GroupMember
	ids := make([]int32, 0)
	err := m.Debug().Select(m.GroupID).Where(m.UserID.Eq(uid)).Scan(&ids)
	if err != nil {
		return Err(c, enum.ErrGroupGetData)
	}

	return Ok(c, ids)
}

// GetGroups 获取群组列表（根据用户位置获取最近的群组列表或者获取最活跃的群组列表）
func (g GroupHandler) GetGroups(c *Ctx) error {
	obj := RetrieveBody(c)
	// page := obj.Page
	// num := obj.Num
	RevertBody(obj)

	return Ok(c, nil)
}

func (g GroupHandler) GetApplyList(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) GetActivitiesByGroupId(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Create(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Apply(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Approve(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Promote(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Transfer(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Remove(c *Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Quit(c *Ctx) error {
	return c.JSON(nil)
}

// put
func (g GroupHandler) Update(c *Ctx) error {
	return c.JSON(nil)
}
