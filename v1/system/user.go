package system

import "github.com/andycai/axe-fiber/v1/hub"

type UserSystem struct {
}

var User = new(UserSystem)

func (u UserSystem) Exists(id uint64) bool {
	_, ok := hub.UserIds[id]

	return ok
}
