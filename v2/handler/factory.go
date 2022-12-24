package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/hub"
	"github.com/spf13/cast"

	"github.com/gofiber/fiber/v2"
)

type Ctx = fiber.Ctx

// 获取远程IP
func IP(c *Ctx) string {
	return c.IP()
}

// RetrieveBody 获取请求体的参数对象
func RetrieveBody(c *Ctx) *comp.BodyObject {
	body := hub.BodyPool.Get().(*comp.BodyObject)
	body.Reset()
	c.BodyParser(body)

	return body
}

// RevertBody 放回对象池
func RevertBody(obj *comp.BodyObject) {
	hub.BodyPool.Put(obj)
}

func String(c *Ctx, key string, defaultValue ...string) string {
	return c.Params(key, defaultValue...)
}

func Int(c *Ctx, key string, defaultValue ...string) int {
	return cast.ToInt(c.Params(key, defaultValue...))
}

func Uint(c *Ctx, key string, defaultValue ...string) uint {
	return cast.ToUint(c.Params(key, defaultValue...))
}

func U32(c *Ctx, key string, defaultValue ...string) uint32 {
	return cast.ToUint32(c.Params(key, defaultValue...))
}

func I32(c *Ctx, key string, defaultValue ...string) int32 {
	return cast.ToInt32(c.Params(key, defaultValue...))
}

func U64(c *Ctx, key string, defaultValue ...string) uint64 {
	return cast.ToUint64(c.Params(key, defaultValue...))
}

func I64(c *Ctx, key string, defaultValue ...string) int64 {
	return cast.ToInt64(c.Params(key, defaultValue...))
}

// 正常响应
func Ok(c *Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": enum.Success,
		"data": data,
	})
}

// 推送响应
func Push(c *Ctx, code int) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  enum.CodeText(code),
	})
}

// 普通消息响应
func Msg(c *Ctx, code int, msg string) error {
	return c.JSON(fiber.Map{
		"code": code,
		"msg":  msg,
	})
}
