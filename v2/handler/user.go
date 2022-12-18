package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/system"
)

type UserHandler struct{}

var User = new(UserHandler)

// 私有方法

// pushUserInfo 推送用户信息给前端
func pushUserInfo(c *Ctx, id int32) error {
	info, err := system.User.GetInfo(id)
	if err != nil {
		return Err(c, enum.ErrUserNotFound)
	}

	return Ok(c, info)
}

// 私有方法 end

// GetCurrentUser 获取当前用户信息
func (uh UserHandler) GetCurrentUser(c *Ctx) error {
	// TODO: 应该从 session 中获取 uid
	var id int32 = 1 // 登录获取的用户ID
	return pushUserInfo(c, id)
}

// GetUser 获取用户信息
func (uh UserHandler) GetUser(c *Ctx) error {
	id := I32(c, "uid")
	return pushUserInfo(c, id)
}

// Login 登录
func (uh UserHandler) Login(c *Ctx) error {
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
	var id int32 = 1 // 登录获取的用户ID
	return pushUserInfo(c, id)
}

// 退出登录
func (uh UserHandler) Exit(c *Ctx) error {
	return Ok(c, nil)
}

func (uh UserHandler) ExitWx(c *Ctx) error {
	return Ok(c, nil)
}

func (uh UserHandler) SaveData(c *Ctx) error {
	return Ok(c, nil)
}
