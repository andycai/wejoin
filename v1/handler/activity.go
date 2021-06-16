package handler

import (
	"github.com/andycai/axe-fiber/define"
	"github.com/andycai/axe-fiber/v1/system"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (a ActivityHandler) GetActivityById(c *fiber.Ctx) error {
	aid := cast.ToUint64(c.Params("aid"))
	if !system.Activity.Exists(aid) {
		return Err(c, define.ErrActivityNotFound)
	}
	act := system.Cache.Activity(aid)

	return Ok(c, act)
}

func (a ActivityHandler) GetActivitiesByUserId(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) GetActivities(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Create(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) End(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Apply(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Cancel(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Remove(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Update(c *fiber.Ctx) error {
	return c.JSON(nil)
}
