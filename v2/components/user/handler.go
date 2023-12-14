package user

import (
	"net/mail"

	"github.com/andycai/axe-fiber/core"
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v2/body"
	"github.com/gofiber/fiber/v2"
)

// pushUserInfo 推送用户信息给前端
func pushUserInfo(c *fiber.Ctx, id int32) error {
	info, err := Dao.GetInfo(id)
	if err != nil {
		return core.Push(c, enum.ErrUserNotFound)
	}

	return core.Ok(c, info)
}

// 私有方法 end

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *fiber.Ctx) error {
	// TODO: 应该从 session 中获取 uid
	var id int32 = 1 // 登录获取的用户ID
	return pushUserInfo(c, id)
}

// GetUser 获取用户信息
func GetUser(c *fiber.Ctx) error {
	id := core.I32(c, "uid")
	return pushUserInfo(c, id)
}

// Register 注册
func Register(c *fiber.Ctx) error {
	username := core.Str(c, "username")
	password := core.Str(c, "password")
	confirmPassword := core.Str(c, "confirmPassword")
	if password != confirmPassword {
		return core.Push(c, enum.ErrTwoPasswordNotMatch)
	}
	err, uid := Dao.Register(username, password, core.IP(c))
	log.Infof("register err: %v, uid: %d", err, uid)
	if err != nil {
		return core.Push(c, enum.ErrUserRegister)
	}
	return core.Ok(c, map[string]int32{"uid": uid})
}

// Login 登录
func Login(c *fiber.Ctx) error {
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
func Exit(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

// ExitWx 退出微信登录
func ExitWx(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

// Update 更新用户数据
func Update(c *fiber.Ctx) error {
	// uid := I32(c, "uid")  校验是否味当前用户
	u := new(body.User)
	if err := c.BodyParser(u); err != nil {
		return core.Push(c, enum.ErrParam)
	}

	// email 校验
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return core.Push(c, enum.ErrUserEmailFormat)
	}

	if err := Dao.Update(u); err != nil {
		return core.Push(c, enum.ErrUserUpdateData)
	}

	return core.Ok(c, nil)
}
