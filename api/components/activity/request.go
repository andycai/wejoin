package activity

import (
	"github.com/andycai/wejoin/model"
	"github.com/gofiber/fiber/v2"
)

type requestCreate struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Kind        int    `json:"kind" form:"kind" validate:"required"`
	Type        int    `json:"type" form:"type" validate:"required"`
	Quota       int    `json:"quota" form:"quota" validate:"required"`
	BeginAt     string `json:"begin_at" form:"begin_at" validate:"required"`
	EndAt       string `json:"end_at" form:"end_at" validate:"required"`
}

func Bind(c *fiber.Ctx, user *model.User) error {
	return nil
}
