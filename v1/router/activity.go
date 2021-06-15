package router

import (
	"github.com/andycai/axe-fiber/v1/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerActivityRouter)
}

func registerActivityRouter(app *fiber.App) {
	actsAPI := app.Group("/activities")
	{
		actsAPI.Get("/:aid", handler.Activity.GetActivityById)
		actsAPI.Get("/", handler.Activity.GetActivities)

		actsAPI.Post("/", handler.Activity.Create)
		actsAPI.Post("/:aid/end", handler.Activity.End)
		actsAPI.Post("/:aid/apply", handler.Activity.Apply)
		actsAPI.Post("/:aid/cancel", handler.Activity.Cancel)
		actsAPI.Post("/:aid/remove/:index", handler.Activity.Remove)

		actsAPI.Put("/:aid}", handler.Activity.Update)
	}
}
