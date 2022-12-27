package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/system"
)

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

func (ah ActivityHandler) GetActivityByID(c *Ctx) error {
	aid := I32(c, "aid")
	info, err := system.Activity.GetInfo(aid)
	if err != nil {
		return Push(c, enum.ErrActivityNotFound)
	}

	return Ok(c, info)
}

func (ah ActivityHandler) GetActivitiesByUserID(c *Ctx) error {
	uid := I32(c, "uid")
	activities, err := system.Activity.GetActivitiesByUserID(uid)
	if err != nil {
		return Push(c, enum.ErrActivityGetData)
	}

	return Ok(c, activities)
}

func (ah ActivityHandler) GetActivities(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Create(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) End(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Apply(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Cancel(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Remove(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) Update(c *Ctx) error {
	return Ok(c, nil)
}
