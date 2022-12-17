package router

import (
	"github.com/andycai/axe-fiber/v2/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	routerNoCheck = make([]func(*fiber.App), 0)
	routerCheck   = make([]func(*fiber.App), 0)
)

func Setup(app *fiber.App) {
	for _, f := range routerCheck {
		f(app)
	}
	for _, f := range routerNoCheck {
		f(app)
	}
	app.Use(middleware.NotFoundRoute)
}
