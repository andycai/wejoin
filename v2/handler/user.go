package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v1/system"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

var User = new(UserHandler)

func (u UserHandler) GetUser(c *fiber.Ctx) error {
	// TODO: 应该从 session 中获取 uid
	uid := U64(c, "uid")
	if !system.User.Exists(uid) {
		return Err(c, enum.ErrUserNotFound)
	}
	user := system.Cache.User(uid)

	return Ok(c, user)
}

func (u UserHandler) Login(c *fiber.Ctx) error {
	// ip := c.Context().RemoteIP().String()

	//err := c.ReadJSON(&b)
	// Checking received data from JSON body.
	/*
		if err := c.BodyParser(signUp); err != nil {
			// Return status 400 and error message.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	*/
	return c.JSON(nil)
}

func (u UserHandler) Exit(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (u UserHandler) ExitWx(c *fiber.Ctx) error {
	return c.JSON(nil)
}

func (u UserHandler) SaveData(c *fiber.Ctx) error {
	return c.JSON(nil)
}
