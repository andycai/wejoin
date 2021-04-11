package router

import (
	"axe/gl"
	"axe/v1/system"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerActivityRouter)
}

func registerActivityRouter(app *fiber.App) {
	actsAPI := fiber.New()
	gl.App.Mount("/activities", actsAPI)
	{
		actsAPI.Get("/:aid", system.Activity.GetActivityById)
		actsAPI.Get("/", system.Activity.GetActivities)

		actsAPI.Post("/", system.Activity.Create)
		actsAPI.Post("/aid/end", system.Activity.End)
		actsAPI.Post("/aid/apply", system.Activity.Apply)
		actsAPI.Post("/aid/cancel", system.Activity.Cancel)
		actsAPI.Post("/aid/remove/:index", system.Activity.Remove)

		actsAPI.Put("/{aid:int64}", system.Activity.Update)
	}
}
