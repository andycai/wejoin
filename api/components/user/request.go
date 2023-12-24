package user

import (
	"github.com/andycai/wejoin/model"
	"github.com/gofiber/fiber/v2"
)

type requestCreate struct {
	ID          uint   `json:"id"`
	Slug        string `json:"slug"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Body        string `json:"body" validate:"required"`
	Action      string `json:"action" form:"action"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	PublishedAt string `json:"published_at" form:"published_at" validate:"required"`
}

func Bind(c *fiber.Ctx, user *model.User) error {
	return nil
}
