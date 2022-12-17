package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerUserRouter)
}

func registerUserRouter(app *fiber.App) {
	app.Post("/login", handler.User.Login)
	app.Post("/login_wx", handler.User.Login)
	app.Post("/login_wx", handler.User.Logout)

	usersAPI := app.Group("/users")
	{
		usersAPI.Get("/:uid", handler.User.GetUser)
		usersAPI.Get("/your/groups", handler.Group.GetGroupsByUserId)
		usersAPI.Get("/your/activities", handler.Activity.GetActivitiesByUserId)
	}
}
