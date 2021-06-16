package hub

import (
	"sync"

	"github.com/andycai/axe-fiber/v1/comp"
)

var BodyPool = &sync.Pool{
	New: func() interface{} {
		return new(comp.BodyObject)
	},
}
