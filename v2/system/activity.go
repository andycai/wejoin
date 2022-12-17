package system

import "github.com/andycai/axe-fiber/v2/hub"

type ActivitySystem struct {
}

var Activity = new(ActivitySystem)

func (a ActivitySystem) Exists(id uint64) bool {
	_, ok := hub.ActivityIDs[id]

	return ok
}
