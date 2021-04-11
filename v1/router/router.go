package router

import (
	"github.com/gofiber/fiber/v2"
)

var (
	routerNoCheck = make([]func(*fiber.App), 0)
)

func InitRouter(app *fiber.App) {
	for _, f := range routerNoCheck {
		f(app)
	}
}
