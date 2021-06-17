package handler

import (
	"github.com/andycai/axe-fiber/define"
	"github.com/andycai/axe-fiber/v1/comp"
	"github.com/andycai/axe-fiber/v1/hub"
	"github.com/spf13/cast"

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

func Ok(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": define.Success,
		"data": data,
	})
}

func GetIP(c *fiber.Ctx) string {
	return c.IP()
}

// RetrieveBody 获取请求体的参数对象
func RetrieveBody(c *fiber.Ctx) *comp.BodyObject {
	body := hub.BodyPool.Get().(*comp.BodyObject)
	body.Reset()
	c.BodyParser(body)

	return body
}

// RevertBody 放回对象池
func RevertBody(obj *comp.BodyObject) {
	hub.BodyPool.Put(obj)
}

func String(c *fiber.Ctx, key string, defaultValue ...string) string {
	return c.Params(key, defaultValue...)
}

func Int(c *fiber.Ctx, key string, defaultValue ...string) int {
	return cast.ToInt(c.Params(key, defaultValue...))
}

func Uint(c *fiber.Ctx, key string, defaultValue ...string) uint {
	return cast.ToUint(c.Params(key, defaultValue...))
}

func U32(c *fiber.Ctx, key string, defaultValue ...string) uint32 {
	return cast.ToUint32(c.Params(key, defaultValue...))
}

func I32(c *fiber.Ctx, key string, defaultValue ...string) int32 {
	return cast.ToInt32(c.Params(key, defaultValue...))
}

func U64(c *fiber.Ctx, key string, defaultValue ...string) uint64 {
	return cast.ToUint64(c.Params(key, defaultValue...))
}

func I64(c *fiber.Ctx, key string, defaultValue ...string) int64 {
	return cast.ToInt64(c.Params(key, defaultValue...))
}

func Err(c *fiber.Ctx, code int) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  define.CodeText(code),
	})
}

func Msg(c *fiber.Ctx, code int, msg string) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
	})
}
