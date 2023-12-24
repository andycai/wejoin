package group

import (
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/enum"
	"github.com/andycai/wejoin/model"
	"github.com/gofiber/fiber/v2"
)

// pushGroupInfo 推送群组信息给前端
func pushGroupInfo(c *fiber.Ctx, gid uint) error {
	info, err := Dao.GetByID(gid)
	if err != nil {
		return core.Push(c, enum.ErrGroupNotFound)
	}

	return core.Ok(c, info)
}

// 私有方法 end

// GetGroupsById 获取单个群组信息
func GetGroupByID(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	return pushGroupInfo(c, gid)
}

// GetGroupsByUserID 获取当前登录用户的群组列表
func GetGroupsByUserID(c *fiber.Ctx) error {
	uid := core.Uint(c, "uid")
	groups, err := Dao.GetByUserID(uid)
	if err != nil {
		return core.Push(c, enum.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// GetGroupsByPage 获取群组列表（根据用户位置获取最近的群组列表或者获取最活跃的群组列表）
func GetGroupsByPage(c *fiber.Ctx) error {
	page := core.Uint(c, "page")
	pageSize := core.Uint(c, "pageSize")

	groups, err := Dao.GetByPage(page, pageSize)
	if err != nil {
		return core.Push(c, enum.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// GetApplicationByGroupID 返回申请加入群组用户列表
func GetApplicationByGroupID(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	list, err := Dao.GetApplictionsByGroupID(gid)
	if err != nil {
		return core.Push(c, enum.ErrGroupApplicationListNotFound)
	}

	return core.Ok(c, list)
}

func Create(c *fiber.Ctx) error {
	// TODO
	Dao.Create(&model.Group{
		//
	})
	return core.Ok(c, nil)
}

func Apply(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	err := Dao.Apply(gid, uid)
	if err != nil {
		return core.Push(c, enum.ErrGroupApply)
	}

	return core.Push(c, enum.SucGroupApply)
}

func Approve(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	mid := core.Uint(c, "mid")
	err := Dao.Approve(gid, uid, mid)
	if err != nil {
		return core.Push(c, enum.ErrGroupApprove)
	}

	return core.Push(c, enum.SucGroupApprove)
}

func Refuse(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	mid := core.Uint(c, "mid")
	err := Dao.Refuse(gid, uid, mid)
	if err != nil {
		return core.Push(c, enum.ErrGroupRefuse)
	}

	return core.Push(c, enum.SucGroupRefuse)
}

func Promote(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	mid := core.Uint(c, "mid")
	err := Dao.Promote(gid, uid, mid)

	if err != nil {
		return core.Push(c, enum.ErrGroupPromote)
	}

	return core.Push(c, enum.SucGroupPromote)
}

func Transfer(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	mid := core.Uint(c, "mid")
	err := Dao.Transfer(gid, uid, mid)

	if err != nil {
		return core.Push(c, enum.ErrGroupTransfer)
	}

	return core.Push(c, enum.SucGroupTransfer)
}

func Fire(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	mid := core.Uint(c, "mid")
	err := Dao.Fire(gid, uid, mid)

	if err != nil {
		return core.Push(c, enum.ErrGroupFire)
	}

	return core.Push(c, enum.SucGroupFire)
}

func Remove(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	err := Dao.Remove(gid, uid)

	if err != nil {
		return core.Push(c, enum.ErrGroupRemove)
	}

	return core.Push(c, enum.SucGroupRemove)
}

func Quit(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	err := Dao.Quit(gid, uid)

	if err != nil {
		return core.Push(c, enum.ErrGroupQuit)
	}

	return core.Push(c, enum.SucGroupQuit)
}

// put
func UpdateName(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	name := core.Str(c, "name")
	// c.Body()
	err := Dao.UpdateName(gid, uid, name)
	if err != nil {
		return core.Push(c, enum.ErrGroupUpdateName)
	}

	return core.Push(c, enum.SucGroupUpdateName)
}

func UpdateNotice(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	notice := core.Str(c, "notice")
	// c.Body()
	err := Dao.UpdateNotice(gid, uid, notice)
	if err != nil {
		return core.Push(c, enum.ErrGroupUpdateNotice)
	}

	return core.Push(c, enum.SucGroupUpdateNotice)
}

func UpdateAddr(c *fiber.Ctx) error {
	gid := core.Uint(c, "gid")
	uid := core.Uint(c, "uid")
	addr := core.Str(c, "addr")
	// c.Body()
	err := Dao.UpdateAddr(gid, uid, addr)
	if err != nil {
		return core.Push(c, enum.ErrGroupUpdateAddr)
	}

	return core.Push(c, enum.SucGroupUpdateAddr)
}
