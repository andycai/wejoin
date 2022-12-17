package hub

import (
	cache2 "github.com/andycai/axe-fiber/library/cache"
)

// LocalUserCache 用户缓存
var LocalUserCache = cache2.New(500)

// LocalGroupCache 群组缓存
var LocalGroupCache = cache2.New(100)

// LocalActivityCache 活动缓存
var LocalActivityCache = cache2.New(1000)
