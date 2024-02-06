package group

import (
	"github.com/andycai/wejoin/model"
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

func BindCreate(c *fiber.Ctx) (*model.Group, error) {
	var r RequestUpdate
	if err := c.BodyParser(&r); err != nil {
		return nil, err
	}

	return &model.Group{
		Scores:  0,
		Level:   1,
		Name:    r.Name,
		Logo:    r.Logo,
		Notice:  r.Notice,
		Address: r.Address,
	}, nil
}

func Bind(c *fiber.Ctx) (*RequestUpdate, error) {
	var r RequestUpdate
	if err := c.BodyParser(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
