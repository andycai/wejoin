package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/system"
	"github.com/spf13/cast"
)

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (a ActivityHandler) GetActivityById(c *Ctx) error {
	aid := cast.ToUint64(c.Params("aid"))
	if !system.Activity.Exists(aid) {
		return Err(c, enum.ErrActivityNotFound)
	}
	act := system.Cache.Activity(aid)

	return Ok(c, act)
}

func (a ActivityHandler) GetActivitiesByUserId(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) GetActivities(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Create(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) End(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Apply(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Cancel(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Remove(c *Ctx) error {
	return c.JSON(nil)
}

func (a ActivityHandler) Update(c *Ctx) error {
	return c.JSON(nil)
}
