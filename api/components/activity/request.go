package activity

import (
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/model"
	"github.com/gofiber/fiber/v2"
)

type RequestCreate struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id" form:"user_id" validate:"required,numeric,min=1,max=999999999"`
	GroupID     uint   `json:"group_id" form:"group_id" validate:"required,numeric,min=0,max=1000"`
	Kind        uint   `json:"kind" form:"kind" validate:"required,numeric,min=1,max=20"`
	Type        uint   `json:"type" form:"type" validate:"required,numeric,min=1,max=5"`
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Quota       uint   `json:"quota" form:"quota" validate:"required,numeric,min=1,max=100"`
	Waiting     uint   `json:"waiting" form:"waiting" validate:"required,numeric,min=0,max=100"`
	Stage       uint   `json:"stage" form:"stage" validate:"required,numeric,min=1,max=4"`
	FeeType     uint   `json:"fee_type" form:"fee_type" validate:"required"`
	FeeMale     uint   `json:"fee_male" form:"fee_male" validate:"required_if=FeeType 2,3,4,5"`
	FeeFemale   uint   `json:"fee_female" form:"fee_female" validate:"required_if=FeeType 2,3,4,5"`
	Address     string `json:"address" form:"address" validate:"required"`
	Ahead       uint   `json:"ahead" form:"ahead" validate:"required,numeric,min=0,max=240"`
	BeginAt     string `json:"begin_at" form:"begin_at" validate:"required,datetime=2006-01-02 15:04:05"`
	EndAt       string `json:"end_at" form:"end_at" validate:"required,datetime=2006-01-02 15:04:05"`
}

func BindCreate(c *fiber.Ctx) (*model.Activity, error) {
	var r RequestCreate
	if err := c.BodyParser(&r); err != nil {
		return nil, err
	}

	if err := core.Validate(&r); err != nil {
		return nil, err
	}

	activity := &model.Activity{
		UserID:      r.UserID,
		GroupID:     r.GroupID,
		Kind:        r.Kind,
		Type:        r.Type,
		Title:       r.Title,
		Description: r.Description,
		Quota:       r.Quota,
		Waiting:     r.Waiting,
		Stage:       r.Stage,
		FeeType:     r.FeeType,
		FeeMale:     r.FeeMale,
		FeeFemale:   r.FeeFemale,
		Address:     r.Address,
		Ahead:       r.Ahead,
		BeginAt:     core.ParseDate(r.BeginAt),
		EndAt:       core.ParseDate(r.EndAt),
	}

	return activity, nil
}

func BindUpdate(c *fiber.Ctx, activity *model.Activity) error {
	var r RequestCreate
	if err := c.BodyParser(&r); err != nil {
		return err
	}

	if err := core.Validate(&r); err != nil {
		return err
	}

	return nil
}
