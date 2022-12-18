package system

import (
	"errors"

	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
)

type UserSystem struct {
}

var User = new(UserSystem)

// GetInfo 获取用户信息
func (us UserSystem) GetInfo(uid int32) (*comp.APIUser, error) {
	u := dao.Q.User
	info := &comp.APIUser{}
	err := u.Where(u.ID.Eq(uid)).Scan(info)
	if info.ID == 0 {
		err = errors.New("not found user data")
	}

	return info, err
}

// Exists 用户是否存在
func (us UserSystem) Exists(uid int32) bool {
	return true
}

// EnterGroup 进入群组
func (us UserSystem) EnterGroup(uid int32, gid int32) {
	//
}

// QuitGroup 离开群组
func (us UserSystem) QuitGroup(uid int32, gid int32) {
	//
}

// ApplyActivity 报名活动
func (us UserSystem) ApplyActivity(uid int32, aid int32) {
	//
}

// CancelActivity 取消报名
func (us UserSystem) CancelActivity(uid int32, aid int32) {
	//
}
