package group

import (
	"github.com/gofiber/fiber/v2"
)

type RequestUpdate struct {
	ID      uint   `json:"id" form:"id"`
	Uid     uint   `json:"uid" form:"uid"`
	Mid     uint   `json:"mid" form:"mid"`
	Name    string `json:"name" form:"name"`
	Logo    string `json:"logo" form:"logo"`
	Notice  string `json:"notice" form:"notice"`
	Address string `json:"address" form:"address"`
}

func Bind(c *fiber.Ctx) (*RequestUpdate, error) {
	var r RequestUpdate
	if err := c.BodyParser(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
