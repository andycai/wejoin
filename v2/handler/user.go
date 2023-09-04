package handler

import (
	"net/mail"

	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v2/body"
	"github.com/andycai/axe-fiber/v2/system"
)

type UserHandler struct{}

var User = new(UserHandler)

// 私有方法

// pushUserInfo 推送用户信息给前端
func pushUserInfo(c *Ctx, id int32) error {
	info, err := system.User.GetInfo(id)
	if err != nil {
		return Push(c, enum.ErrUserNotFound)
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

// Register 注册
func (uh UserHandler) Register(c *Ctx) error {
	username := Str(c, "username")
	password := Str(c, "password")
	confirmPassword := Str(c, "confirmPassword")
	if password != confirmPassword {
		return Push(c, enum.ErrTwoPasswordNotMatch)
	}
	err, uid := system.User.Register(username, password, IP(c))
	log.Infof("register err: %v, uid: %d", err, uid)
	if err != nil {
		return Push(c, enum.ErrUserRegister)
	}
	return Ok(c, map[string]int32{"uid": uid})
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

// Exit 退出登录
func (uh UserHandler) Exit(c *Ctx) error {
	return Ok(c, nil)
}

// ExitWx 退出微信登录
func (uh UserHandler) ExitWx(c *Ctx) error {
	return Ok(c, nil)
}

// Update 更新用户数据
func (uh UserHandler) Update(c *Ctx) error {
	// uid := I32(c, "uid")  校验是否味当前用户
	u := new(body.User)
	if err := c.BodyParser(u); err != nil {
		return Push(c, enum.ErrParam)
	}

	// email 校验
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return Push(c, enum.ErrUserEmailFormat)
	}

	if err := system.User.Update(u); err != nil {
		return Push(c, enum.ErrUserUpdateData)
	}

	return Ok(c, nil)
}
