package handler

import (
	"github.com/andycai/axe-fiber/v1/system"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type UserHandler struct{}

var User = new(UserHandler)

func (u UserHandler) GetUser(ctx *fiber.Ctx) error {
	uid := cast.ToUint64(ctx.Params("uid"))
	user := system.Cache.User(uid)

	return Ok(ctx, user)
}

func (u UserHandler) Login(ctx *fiber.Ctx) error {
	//err := ctx.ReadJSON(&b)
	// Checking received data from JSON body.
	/*
		if err := ctx.BodyParser(signUp); err != nil {
			// Return status 400 and error message.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	*/
	return ctx.JSON(nil)
}

func (u UserHandler) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserHandler) EnterGroup(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserHandler) QuitGroup(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserHandler) ApplyActivity(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserHandler) Cancel(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserHandler) SaveData(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
