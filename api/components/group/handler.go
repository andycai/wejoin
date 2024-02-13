package group

import (
	"github.com/andycai/wejoin/constant"
	"github.com/andycai/wejoin/core"
	"github.com/gofiber/fiber/v2"
)

// pushGroupInfo 推送群组信息给前端
func pushGroupInfo(c *fiber.Ctx, gid uint) error {
	info, err := Dao.GetByID(gid)
	if err != nil {
		return core.Push(c, constant.ErrGroupNotFound)
	}

	return core.Ok(c, info)
}

// 私有方法 end

// GetGroupsById 获取单个群组信息
func handleGetGroupByID(c *fiber.Ctx) error {
	gid := core.Uint(c, "id")
	return pushGroupInfo(c, gid)
}

// GetGroupsByUserID 获取当前登录用户的群组列表
func GetGroupsByUserID(c *fiber.Ctx) error {
	uid := core.Uint(c, "uid")
	groups, err := Dao.GetByUserID(uid)
	if err != nil {
		return core.Push(c, constant.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// handleGetGroupsByPage 获取群组列表（根据用户位置获取最近的群组列表或者获取最活跃的群组列表）
func handleGetGroupsByPage(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", constant.PAGE_SIZE)

	groups, err := Dao.GetByPage(page, pageSize)
	if err != nil {
		return core.Push(c, constant.ErrGroupGetData)
	}

	return core.Ok(c, groups)
}

// handleGetApplicationByGroupID 返回申请加入群组用户列表
func handleGetApplicationByGroupID(c *fiber.Ctx) error {
	gid := core.Uint(c, "id")
	list, err := Dao.GetApplictionsByGroupID(gid)
	if err != nil {
		return core.Push(c, constant.ErrGroupApplicationListNotFound)
	}

	return core.Ok(c, list)
}

func handleCreate(c *fiber.Ctx) error {
	groupVo, err := BindCreate(c)
	if err != nil {
		return core.Err(c, err)
	}
	Dao.Create(groupVo)
	return core.Ok(c, nil)
}

func handleApply(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Apply(r.ID, r.Uid)
	if err != nil {
		return core.Push(c, constant.ErrGroupApply)
	}

	return core.Push(c, constant.SucGroupApply)
}

func handleCancel(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Cancel(r.ID, r.Uid)
	if err != nil {
		return core.Push(c, constant.ErrGroupCancel)
	}

	return core.Push(c, constant.SucGroupCancel)
}

func handleApprove(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Approve(r.ID, r.Uid, r.Mid)
	if err != nil {
		return core.Push(c, constant.ErrGroupApprove)
	}

	return core.Push(c, constant.SucGroupApprove)
}

func handleRefuse(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Refuse(r.ID, r.Uid, r.Mid)
	if err != nil {
		return core.Push(c, constant.ErrGroupRefuse)
	}

	return core.Push(c, constant.SucGroupRefuse)
}

func handlePromote(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Promote(r.ID, r.Uid, r.Mid)

	if err != nil {
		return core.Push(c, constant.ErrGroupPromote)
	}

	return core.Push(c, constant.SucGroupPromote)
}

func handleTransfer(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Transfer(r.ID, r.Uid, r.Mid)

	if err != nil {
		return core.Push(c, constant.ErrGroupTransfer)
	}

	return core.Push(c, constant.SucGroupTransfer)
}

func handleFire(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.Fire(r.ID, r.Uid, r.Mid)

	if err != nil {
		return core.Push(c, constant.ErrGroupFire)
	}

	return core.Push(c, constant.SucGroupFire)
}

func handleRemove(c *fiber.Ctx) error {
	var r RequestUpdate
	if err := c.BodyParser(&r); err != nil {
		return err
	}
	err := Dao.Remove(r.ID, r.Uid)

	if err != nil {
		return core.Push(c, constant.ErrGroupRemove)
	}

	return core.Push(c, constant.SucGroupRemove)
}

func handleQuit(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}

	err = Dao.Quit(r.ID, r.Mid)
	if err != nil {
		return core.Push(c, constant.ErrGroupQuit)
	}

	return core.Push(c, constant.SucGroupQuit)
}

// put
func handleUpdateName(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.UpdateName(r.ID, r.Uid, r.Name)
	if err != nil {
		return core.Push(c, constant.ErrGroupUpdateName)
	}

	return core.Push(c, constant.SucGroupUpdateName)
}

func handleUpdateLogo(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.UpdateLogo(r.ID, r.Uid, r.Logo)
	if err != nil {
		return core.Push(c, constant.ErrGroupUpdateName)
	}

	return core.Push(c, constant.SucGroupUpdateName)
}

func handleUpdateNotice(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.UpdateNotice(r.ID, r.Uid, r.Notice)
	if err != nil {
		return core.Push(c, constant.ErrGroupUpdateNotice)
	}

	return core.Push(c, constant.SucGroupUpdateNotice)
}

func handleUpdateAddress(c *fiber.Ctx) error {
	r, err := Bind(c)
	if err != nil {
		return core.Err(c, err)
	}
	err = Dao.UpdateAddress(r.ID, r.Uid, r.Address)
	if err != nil {
		return core.Push(c, constant.ErrGroupUpdateAddr)
	}

	return core.Push(c, constant.SucGroupUpdateAddr)
}
