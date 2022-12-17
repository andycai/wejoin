package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v1/system"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type GroupHandler struct{}

var Group = new(GroupHandler)

// GetGroupsById 获取单个群组信息
func (g GroupHandler) GetGroupById(c *fiber.Ctx) error {
	gid := cast.ToUint64(c.Params("gid"))
	if !system.Group.Exists(gid) {
		return Err(c, enum.ErrGroupNotFound)
	}
	group := system.Cache.Group(gid)

	return Ok(c, group)
}

// GetGroupsByUserId 获取当前登录用户的群组列表
func (g GroupHandler) GetGroupsByUserId(c *fiber.Ctx) error {
	return c.JSON(nil)
}

// GetGroups 获取群组列表（根据用户位置获取最近的群组列表或者获取最活跃的群组列表）
func (g GroupHandler) GetGroups(c *fiber.Ctx) error {
	obj := RetrieveBody(c)
	// page := obj.Page
	// num := obj.Num
	RevertBody(obj)

	return Ok(c, nil)
}

func (g GroupHandler) GetApplyList(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) GetActivitiesByGroupId(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Create(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Apply(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Approve(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Promote(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Transfer(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Remove(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (g GroupHandler) Quit(c *fiber.Ctx) error {
	return c.JSON(nil)
}

// put
func (g GroupHandler) Update(c *fiber.Ctx) error {
	return c.JSON(nil)
}
