package system

import "github.com/andycai/axe-fiber/v2/hub"

type GroupSystem struct {
}

var Group = new(GroupSystem)

func (u GroupSystem) Exists(id uint64) bool {
	_, ok := hub.GroupIDs[id]

	return ok
}
