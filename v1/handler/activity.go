package handler

import "github.com/gofiber/fiber/v2"

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (a ActivityHandler) GetActivityById(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
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
