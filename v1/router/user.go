package router

import (
	"axe/gl"
	"axe/v1/system"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerUserRouter)
}

func registerUserRouter(app *fiber.App) {
	app.Post("/login", system.User.Login)
	app.Post("/login_wx", system.User.Login)
	app.Post("/login_wx", system.User.Logout)

	usersAPI := fiber.New()
	gl.App.Mount("/users", usersAPI)
	{
		usersAPI.Get("/{uid:int64}", system.User.GetUser)
		usersAPI.Get("/your/groups", system.Group.GetGroupsByUserId)
		usersAPI.Get("/your/activities", system.Activity.GetActivitiesByUserId)
	}
}
