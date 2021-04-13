package middleware

import (
	"github.com/andycai/axe-fiber/gl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		// Log each request
		gl.Log.Info(
			"fetch URL",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
		)

		// Go to next middleware
		return c.Next()
	})

	app.Use(
		recover.New(),
		cors.New(),
		requestid.New(),
		logger.New(logger.Config{
			Format:     "${time} ${pid} ${locals:requestid} ${status} - ${method} ${path}​\n​",
			TimeFormat: "2006-01-02 15:04:05",
			// TimeZone:   "America/New_York",
		}),
	)
}
