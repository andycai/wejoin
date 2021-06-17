package cache

import (
	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
	"go.uber.org/zap"
)

type UserCache struct {
	usersForName map[string]*comp.User
	usersForId   map[uint64]*comp.User
	sessions     map[string]*comp.Session
}

var User = &UserCache{
	usersForName: make(map[string]*comp.User),
	usersForId:   make(map[uint64]*comp.User),
	sessions:     make(map[string]*comp.Session),
}

func (uc *UserCache) GetUserById(id uint64) *comp.User {
	if user, ok := uc.usersForId[id]; ok {
		log.Info("获得用户数据（缓存）：", zap.Uint64("uid", id))
		return user
	}
	log.Info("获得用户数据（DB）：", zap.Uint64("uid", id))
	user := mysql.User.GetUserById(id)
	user.OutDB()
	uc.usersForId[id] = user

	return user
}
