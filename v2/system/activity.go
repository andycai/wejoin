package system

type ActivitySystem struct {
}

var Activity = new(ActivitySystem)

func (a ActivitySystem) Exists(id int32) bool {
	return true
}
