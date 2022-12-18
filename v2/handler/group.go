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
		return Err(c, enum.ErrGroupNotFound)
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
		return Err(c, enum.ErrGroupGetData)
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
		return Err(c, enum.ErrGroupGetData)
	}

	return Ok(c, groups)
}

func (gh GroupHandler) GetApplyList(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) GetActivitiesByGroupId(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Create(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Apply(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Approve(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Promote(c *Ctx) error {
	return Ok(c, nil)
}

func (gh GroupHandler) Transfer(c *Ctx) error {
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
