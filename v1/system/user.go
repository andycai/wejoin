package system

import (
	"log"

	"github.com/andycai/axe-fiber/cache"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
)

type UserSystem struct{}

var User = new(UserSystem)

func (u UserSystem) GetUser(ctx *fiber.Ctx) error {
	uid := cast.ToInt64(ctx.Params("uid"))
	user := cache.User.GetUserById(uid)
	// ctx.Context().Logger().Printf("user: %d", uid)
	log.Printf("user: %d", uid)

	return Ok(ctx, user)
}

func (u UserSystem) Login(ctx *fiber.Ctx) error {
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

func (u UserSystem) Logout(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserSystem) EnterGroup(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserSystem) QuitGroup(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserSystem) ApplyActivity(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserSystem) Cancel(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (u UserSystem) SaveData(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
