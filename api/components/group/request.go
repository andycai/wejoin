package group

import (
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/model"
	"github.com/gofiber/fiber/v2"
)

type RequestUpdate struct {
	ID      uint   `json:"id"`
	Uid     uint   `json:"uid"`
	Name    string `json:"name" form:"name" validate:"required"`
	Logo    string `json:"logo" form:"logo" validate:"required"`
	Notice  string `json:"notice" form:"notice" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}

func Bind(c *fiber.Ctx, user *model.Group) error {
	var r RequestUpdate
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if err := core.Validate(r); err != nil {
		return err
	}

	return nil
}
