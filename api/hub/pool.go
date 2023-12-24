package hub

import (
	"sync"

	"github.com/andycai/wejoin/api/comp"
)

var BodyPool = &sync.Pool{
	New: func() interface{} {
		return new(comp.BodyObject)
	},
}
