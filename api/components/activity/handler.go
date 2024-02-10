package activity

import (
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/enum"
	"github.com/gofiber/fiber/v2"
)

// handleGetByID get the activity by ID
func handleGetByID(c *fiber.Ctx) error {
	aid := core.Uint(c, "id")
	info, err := Dao.GetByID(aid)
	if err != nil {
		return core.Push(c, enum.ErrActivityNotFound)
	}

	return core.Ok(c, info)
}

// handleGetByUserID get the activities of the user
func handleGetByUserID(c *fiber.Ctx) error {
	uid := core.Uint(c, "uid")
	activities, err := Dao.GetByUserID(uid)
	if err != nil {
		return core.Push(c, enum.ErrActivityGetData)
	}

	return core.Ok(c, activities)
}

// handleGetByOrganizerUserID get the activities of the organizer
func handleGetByOrganizerUserID(c *fiber.Ctx) error {
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

// handleGetByPage get the activities by page
func handleGetByPage(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", enum.PAGE_SIZE)

	groups, err := Dao.GetByPage(page, pageSize)
	if err != nil {
		return core.Push(c, enum.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// handleGetByGroupID get the activities by the group ID
func handleGetByGroupID(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	list, err := Dao.GetByGroupID(gid)
	if err != nil {
		return core.Push(c, enum.ErrActivityGetData)
	}

	return core.Ok(c, list)
}

func handleCreate(c *fiber.Ctx) error {
	activity, err := BindCreate(c)
	if err != nil {
		return core.Push(c, enum.ErrActivityCreate)
	}

	if err := Dao.Create(activity); err != nil {
		return core.Push(c, enum.ErrActivityCreate)
	}

	return core.Ok(c, nil)
}

func handleComplete(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func handleEnd(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func handleApply(c *fiber.Ctx) error {
	activityUser, err := BindCreateActivityUser(c)
	if err != nil {
		return core.Push(c, enum.ErrActivityApply)
	}
	if err := Dao.CreateActivityUser(activityUser); err != nil {
		return core.Push(c, enum.ErrActivityApply)
	}

	return core.Ok(c, nil)
}

func handleCancel(c *fiber.Ctx) error {
	return core.Ok(c, nil)
}

func handleRemove(c *fiber.Ctx) error {
	// activity := new(body.Activity)
	// if err := c.BodyParser(activity); err != nil {
	// 	return core.Push(c, enum.ErrParam)
	// }
	// if err := Dao.Update(activity); err != nil {
	// 	return core.Push(c, enum.ErrActivityUpdate)
	// }
	return core.Ok(c, nil)
}

func handleUpdate(c *fiber.Ctx) error {
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
