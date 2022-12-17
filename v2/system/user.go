package system

import "github.com/andycai/axe-fiber/v2/hub"

type UserSystem struct {
}

var User = new(UserSystem)

// Exists 用户是否存在
func (u UserSystem) Exists(id uint64) bool {
	_, ok := hub.UserIds[id]

	return ok
}

// EnterGroup 进入群组
func (u UserSystem) EnterGroup(uid uint64, gid uint64) {
	//
}

// QuitGroup 离开群组
func (u UserSystem) QuitGroup(uid uint64, gid uint64) {
	//
}

// ApplyActivity 报名活动
func (u UserSystem) ApplyActivity(uid uint64, aid uint64) {
	//
}

// CancelActivity 取消报名
func (u UserSystem) CancelActivity(uid uint64, aid uint64) {
	//
}
