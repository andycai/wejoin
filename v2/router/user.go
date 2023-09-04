package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerUserCheckRouter)
	routerNoCheck = append(routerNoCheck, registerUserNoCheckRouter)
}

func registerUserNoCheckRouter(r fiber.Router) {
	// 注册
	r.Post("/register", handler.User.Register)
	// 登录
	r.Post("/login", handler.User.Login)
	// 退出
	r.Post("/exit", handler.User.Exit)
	// 微信登录
	r.Post("/login_wx", handler.User.Login)
	r.Post("/exit_wx", handler.User.ExitWx)
}

func registerUserCheckRouter(r fiber.Router) {
	api := r.Group("/users")
	{
		api.Get("/:uid", handler.User.GetUser)
		api.Put("/:uid", handler.User.Update)
	}
}
