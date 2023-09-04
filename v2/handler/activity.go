package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/body"
	"github.com/andycai/axe-fiber/v2/system"
)

type ActivityHandler struct{}

var Activity = new(ActivityHandler)

// GetActivityByID 返回活动详情
func (ah ActivityHandler) GetActivityByID(c *Ctx) error {
	aid := I32(c, "aid")
	info, err := system.Activity.GetInfo(aid)
	if err != nil {
		return Push(c, enum.ErrActivityNotFound)
	}

	return Ok(c, info)
}

// GetActivitiesByGroupID 返回用户参与的活动列表
func (ah ActivityHandler) GetActivitiesByUserID(c *Ctx) error {
	uid := I32(c, "uid")
	activities, err := system.Activity.GetActivitiesByUserID(uid)
	if err != nil {
		return Push(c, enum.ErrActivityGetData)
	}

	return Ok(c, activities)
}

// GetActivitiesByGroupID 返回群组创建的活动列表
// func (ah ActivityHandler) GetActivitiesByGroupID(c *Ctx) error {
// 	gid := I32(c, "gid")
// 	activities, err := system.Activity.GetActivitiesByGroupID(gid)
// 	if err != nil {
// 		return Push(c, enum.ErrActivityGetData)
// 	}

// 	return Ok(c, activities)
// }

// GetActivities 返回活动列表
func (ah ActivityHandler) GetActivities(c *Ctx) error {
	page := Int(c, "page")
	num := Int(c, "num")

	groups, err := system.Activity.GetActivities(page, num)
	if err != nil {
		return Push(c, enum.ErrGroupGetData)
	}

	return Ok(c, groups)
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
	activity := new(body.Activity)
	if err := c.BodyParser(activity); err != nil {
		return Push(c, enum.ErrParam)
	}
	if err := system.Activity.Update(activity); err != nil {
		return Push(c, enum.ErrActivityUpdate)
	}
	return Ok(c, nil)
}

func (ah ActivityHandler) Update(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateTitle(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateContent(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateAddr(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateTime(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateQuota(c *Ctx) error {
	return Ok(c, nil)
}

func (ah ActivityHandler) UpdateFee(c *Ctx) error {
	return Ok(c, nil)
}
