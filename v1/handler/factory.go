package handler

import (
	"github.com/andycai/axe-fiber/define"

	"github.com/gofiber/fiber/v2"
)

type Context struct {
	Ctx *fiber.Ctx
}

func (c *Context) JSON(data interface{}) error {
	return c.Ctx.JSON(data)
}

func (c *Context) Params(key string, defaultValue ...string) string {
	return c.Ctx.Params(key, defaultValue...)
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(fiber.Map{
		"code": define.Success,
		"data": data,
	})
}

func Err(ctx *fiber.Ctx, code int) error {
	return ctx.JSON(fiber.Map{
		"code": code,
		"msg":  define.CodeText(code),
	})
}

func Msg(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
	})
}
