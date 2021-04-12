package system

import "github.com/gofiber/fiber/v2"

type ActivitySystem struct{}

var Activity = new(ActivitySystem)

func (a ActivitySystem) GetActivityById(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) GetActivitiesByUserId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) GetActivities(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) End(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) Apply(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) Cancel(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) Remove(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (a ActivitySystem) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
