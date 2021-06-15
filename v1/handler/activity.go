package handler

import (
	"github.com/andycai/axe-fiber/define"
	"github.com/andycai/axe-fiber/v1/system"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (a ActivityHandler) GetActivityById(ctx *fiber.Ctx) error {
	aid := cast.ToUint64(ctx.Params("aid"))
	if !system.Activity.Exists(aid) {
		return Err(ctx, define.ErrActivityNotFound)
	}
	act := system.Cache.Activity(aid)

	return Ok(ctx, act)
}

func (a ActivityHandler) GetActivitiesByUserId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) GetActivities(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) End(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) Apply(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) Cancel(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) Remove(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivityHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
