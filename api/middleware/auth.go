package middleware

import (
	"github.com/andycai/wejoin/library/authentication"
	"github.com/gofiber/fiber/v2"
)

func Authorize(c *fiber.Ctx) error {
	isAuthenticated, _ := authentication.AuthGet(c)

	if isAuthenticated {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg":   "Unauthorized",
	})
}
