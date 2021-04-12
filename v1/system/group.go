package system

import "github.com/gofiber/fiber/v2"

type GroupSystem struct{}

var Group = new(GroupSystem)

func (g GroupSystem) GetGroupById(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) GetGroupsByUserId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) GetGroups(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) GetApplyList(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) GetActivitiesByGroupId(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Create(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Apply(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Approve(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Promote(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Transfer(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Remove(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (g GroupSystem) Quit(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

// put
func (g GroupSystem) Update(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
