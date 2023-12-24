package activity

import (
	"github.com/andycai/wejoin/api/body"
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/enum"
	"github.com/gofiber/fiber/v2"
)

// GetActivityByID 返回活动详情
func GetActivityByID(c *fiber.Ctx) error {
	aid := core.I32(c, "aid")
	info, err := Dao.GetInfo(aid)
	if err != nil {
		return core.Push(c, enum.ErrActivityNotFound)
	}

	return core.Ok(c, info)
}

// GetActivitiesByGroupID 返回用户参与的活动列表
func GetActivitiesByUserID(c *fiber.Ctx) error {
	uid := core.I32(c, "uid")
	activities, err := Dao.GetActivitiesByUserID(uid)
	if err != nil {
		return core.Push(c, enum.ErrActivityGetData)
	}

	return core.Ok(c, activities)
}

// GetActivitiesByGroupID 返回群组创建的活动列表
// func GetActivitiesByGroupID(c *fiber.Ctx) error {
// 	gid := I32(c, "gid")
// 	activities, err := Dao.GetActivitiesByGroupID(gid)
// 	if err != nil {
// 		return Push(c, enum.ErrActivityGetData)
// 	}

// 	return Ok(c, activities)
// }

// GetActivitiesByPage 返回活动列表
func GetActivitiesByPage(c *fiber.Ctx) error {
	page := core.Int(c, "page")
	num := core.Int(c, "num")

	groups, err := Dao.GetActivities(page, num)
	if err != nil {
		return core.Push(c, enum.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// GetActivitiesByGroupId 返回群组创建活动列表
func GetActivitiesByGroupId(c *fiber.Ctx) error {
	gid := core.I32(c, "gid")
	list, err := Dao.GetActivitiesByGroupID(gid)
	if err != nil {
		return core.Push(c, enum.ErrActivityGetData)
	}

	return core.Ok(c, list)
}

func Create(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func End(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func Apply(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func Cancel(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func Remove(c *fiber.Ctx) error {
	activity := new(body.Activity)
	if err := c.BodyParser(activity); err != nil {
		return core.Push(c, enum.ErrParam)
	}
	if err := Dao.Update(activity); err != nil {
		return core.Push(c, enum.ErrActivityUpdate)
	}
	return core.Ok(c, nil)
}

func Update(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateTitle(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateContent(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateAddr(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateTime(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateQuota(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func UpdateFee(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}
