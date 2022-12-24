package handler

import (
	"github.com/andycai/axe-fiber/enum"
	"github.com/andycai/axe-fiber/v2/system"
)

type GroupHandler struct{}

var Group = new(GroupHandler)

// 私有方法

// pushGroupInfo 推送群组信息给前端
func pushGroupInfo(c *Ctx, id int32) error {
	info, err := system.Group.GetInfo(id)
	if err != nil {
		return Push(c, enum.ErrGroupNotFound)
	}

	return Ok(c, info)
}

// 私有方法 end

// GetGroupsById 获取单个群组信息
func (gh GroupHandler) GetGroupByID(c *Ctx) error {
	gid := I32(c, "gid")
	return pushGroupInfo(c, gid)
}

// GetGroupsByUserID 获取当前登录用户的群组列表
func (gh GroupHandler) GetGroupsByUserID(c *Ctx) error {
	uid := I32(c, "uid")
	groups, err := system.Group.GetGroupsByUserID(uid)
	if err != nil {
		return Push(c, enum.ErrGroupGetData)
	}

	return Ok(c, groups)
}

// GetGroups 获取群组列表（根据用户位置获取最近的群组列表或者获取最活跃的群组列表）
func (gh GroupHandler) GetGroups(c *Ctx) error {
	obj := RetrieveBody(c)
	page := Int(c, "page")
	num := Int(c, "num")
	RevertBody(obj)

	groups, err := system.Group.GetGroups(page, num)
	if err != nil {
		return Push(c, enum.ErrGroupGetData)
	}

	return Ok(c, groups)
}

func (gh GroupHandler) GetApplyList(c *Ctx) error {
	gid := I32(c, "gid")
	list, err := system.Group.GetApplyList(gid)
	if err != nil {
		return Push(c, enum.ErrGroupApplicationListNotFound)
	}

	return Ok(c, list)
}

func (gh GroupHandler) GetActivitiesByGroupId(c *Ctx) error {
	gid := I32(c, "gid")
	list, err := system.Activity.GetActivitiesByGroupID(gid)
	if err != nil {
		return Push(c, enum.ErrActivityGetData)
	}

	return Ok(c, list)
}

func (gh GroupHandler) Create(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Apply(c *Ctx) error {
	gid := I32(c, "gid")
	uid := I32(c, "uid")
	err := system.Group.Apply(gid, uid)
	if err != nil {
		return Push(c, enum.ErrGroupApply)
	}

	return Push(c, enum.SucGroupApply)
}

func (gh GroupHandler) Approve(c *Ctx) error {
	gid := I32(c, "gid")
	uid := I32(c, "uid")
	err := system.Group.Approve(gid, uid)
	if err != nil {
		return Push(c, enum.ErrGroupApprove)
	}

	return Push(c, enum.SucGroupApprove)
}

func (gh GroupHandler) Refuse(c *Ctx) error {
	gid := I32(c, "gid")
	uid := I32(c, "uid")
	err := system.Group.Refuse(gid, uid)
	if err != nil {
		return Push(c, enum.ErrGroupRefuse)
	}

	return Push(c, enum.SucGroupRefuse)
}

func (gh GroupHandler) Promote(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Transfer(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Fire(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Remove(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Quit(c *Ctx) error {
	return Ok(c, nil)
}

// put
func (gh GroupHandler) Update(c *Ctx) error {
	return Ok(c, nil)
}
