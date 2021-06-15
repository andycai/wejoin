package handler

import "github.com/gofiber/fiber/v2"

type GroupHandler struct{}

var Group = new(GroupHandler)

func (g GroupHandler) GetGroupById(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) GetGroupsByUserId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) GetGroups(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) GetApplyList(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) GetActivitiesByGroupId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Apply(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Approve(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Promote(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Transfer(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Remove(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupHandler) Quit(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

// put
func (g GroupHandler) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
