package system

import (
	"fmt"

	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
	"github.com/andycai/axe-fiber/v1/hub"
)

// userCacheKey 用户缓存 key
func userCacheKey(id uint64) string {
	return fmt.Sprintf("User:%d", id)
}

// InitUserIds 缓存用户 id
func (c CacheSystem) InitUserIds() {
	userIDs := mysql.User.GetIds()
	for _, v := range userIDs {
		hub.UserIds[v.ID] = struct{}{}
	}
	groupIDs := mysql.Group.GetIds()
	for _, v := range groupIDs {
		hub.GroupIDs[v.ID] = struct{}{}
	}
	activityIDs := mysql.Activity.GetIds()
	for _, v := range activityIDs {
		hub.ActivityIDs[v.ID] = struct{}{}
	}
}

// ExistsUser 是否存在用户
func (c CacheSystem) ExistsUser(id uint64) bool {
	return hub.LocalUserCache.Has(userCacheKey(id))
}

func (c CacheSystem) User(id uint64) *comp.User {
	if !c.ExistsUser(id) {
		log.Infof("从 DB 获取用户数据，id:%d", id)
		user := mysql.User.GetUserById(id)
		user.OutDB()
		err := hub.LocalUserCache.Set(userCacheKey(id), user)
		if err != nil {
			log.Errorf("从 DB 获取用户数据失败，id:%d", id)
			return nil
		}
	}

	user, err := hub.LocalUserCache.Get(userCacheKey(id))
	if err != nil {
		log.Errorf("获取用户数据失败，id:%d", id)
		return nil
	}
	return user.(*comp.User)
}
