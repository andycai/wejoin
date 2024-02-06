package activity

import (
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/enum"
	"github.com/gofiber/fiber/v2"
)

// GetByID get the activity by ID
func GetByID(c *fiber.Ctx) error {
	aid := core.Uint(c, "id")
	info, err := Dao.GetByID(aid)
	if err != nil {
		return core.Push(c, enum.ErrActivityNotFound)
	}

	return core.Ok(c, info)
}

// GetByUserID get the activities of the user
func GetByUserID(c *fiber.Ctx) error {
	uid := core.Uint(c, "uid")
	activities, err := Dao.GetByUserID(uid)
	if err != nil {
		return core.Push(c, enum.ErrActivityGetData)
	}

	return core.Ok(c, activities)
}

// GetByOrganizerUserID get the activities of the organizer
func GetByOrganizerUserID(c *fiber.Ctx) error {
	uid := core.Uint(c, "uid")
	activities, err := Dao.GetByOrganizerUserID(uid)
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

// GetByPage get the activities by page
func GetByPage(c *fiber.Ctx) error {
	page := c.QueryInt("page")
	pageSize := c.QueryInt("page_size")

	groups, err := Dao.GetByPage(page, pageSize)
	if err != nil {
		return core.Push(c, enum.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// GetByGroupID get the activities by the group ID
func GetByGroupID(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	list, err := Dao.GetByGroupID(gid)
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
	// activity := new(body.Activity)
	// if err := c.BodyParser(activity); err != nil {
	// 	return core.Push(c, enum.ErrParam)
	// }
	// if err := Dao.Update(activity); err != nil {
	// 	return core.Push(c, enum.ErrActivityUpdate)
	// }
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
