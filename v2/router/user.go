package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerUserRouter)
}

func registerUserRouter(app *fiber.App) {
	// v2 版本路由
	// 登录
	app.Post("/v2/login", handler.User.Login)
	// 退出
	app.Post("/v2/exit", handler.User.Exit)
	// 微信登录
	app.Post("/v2/login_wx", handler.User.Login)
	app.Post("/v2/exit_wx", handler.User.ExitWx)

	usersAPI := app.Group("/v2/users")
	{
		usersAPI.Get("/:uid", handler.User.GetUser)
		usersAPI.Get("/your/groups", handler.Group.GetGroupsByUserId)
		usersAPI.Get("/your/activities", handler.Activity.GetActivitiesByUserId)
	}
}
