package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/dao"
)

type UserHandler struct{}

var User = new(UserHandler)

// GetCurrentUser 获取当前用户信息
func (u UserHandler) GetCurrentUser(c *Ctx) error {
	// TODO: 应该从 session 中获取 uid
	return Err(c, enum.ErrUserNotFound)
}

// GetUser 获取用户信息
func (u UserHandler) GetUser(c *Ctx) error {
	id := I32(c, "uid")
	q := dao.Q.User
	// user, err := q.Debug().Where(q.ID.Eq(uid)).Take()
	user, err := q.Where(q.ID.Eq(id)).Take()
	if err != nil {
		return Err(c, enum.ErrUserNotFound)
	}

	return Ok(c, user)
}

// Login 登录
func (u UserHandler) Login(c *Ctx) error {
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

// 退出登录
func (u UserHandler) Exit(c *Ctx) error {
	return c.JSON(nil)
}

func (u UserHandler) ExitWx(c *Ctx) error {
	return c.JSON(nil)
}

func (u UserHandler) SaveData(c *Ctx) error {
	return c.JSON(nil)
}
