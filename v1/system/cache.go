package system

import (
	"fmt"

	"github.com/andycai/axe-fiber/dao/mysql"

	"github.com/andycai/axe-fiber/log"

	"github.com/andycai/axe-fiber/v1/comp"
	"github.com/andycai/axe-fiber/v1/hub"
)

type CacheSystem struct {
}

var Cache = new(CacheSystem)

// begin 私有方法

// UserCacheKey 用户缓存 key
func UserCacheKey(id uint64) string {
	return fmt.Sprintf("User:%d", id)
}

// GroupCacheKey 群组缓存 key
func GroupCacheKey(id uint64) string {
	return fmt.Sprintf("Group:%d", id)
}

// ActivityCacheKey 活动缓存 key
func ActivityCacheKey(id uint64) string {
	return fmt.Sprintf("Activity:%d", id)
}

// end 私有方法

// InitIds 缓存用户、群组和活动的 id
func (c CacheSystem) InitIds() {
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
	return hub.LocalUserCache.Has(UserCacheKey(id))
}

func (c CacheSystem) User(id uint64) *comp.User {
	if !c.ExistsUser(id) {
		log.Infof("从 DB 获取用户数据，id:%d", id)
		user := mysql.User.GetUserById(id)
		user.OutDB()
		err := hub.LocalUserCache.Set(UserCacheKey(id), user)
		if err != nil {
			log.Errorf("从 DB 获取用户数据失败，id:%d", id)
			return nil
		}
	}

	user, err := hub.LocalUserCache.Get(UserCacheKey(id))
	if err != nil {
		log.Errorf("获取用户数据失败，id:%d", id)
		return nil
	}
	return user.(*comp.User)
}

// ExistsGroup 是否存在群组
func (c CacheSystem) ExistsGroup(id uint64) bool {
	return hub.LocalGroupCache.Has(GroupCacheKey(id))
}

func (c CacheSystem) Group(id uint64) *comp.Group {
	if !c.ExistsGroup(id) {
		log.Infof("从 DB 获取群组数据，id:%d", id)
		group := mysql.Group.GetGroupById(id)
		group.OutDB()
		err := hub.LocalGroupCache.Set(GroupCacheKey(id), group)
		if err != nil {
			log.Errorf("从 DB 获取群组数据失败，id:%d", id)
			return nil
		}
	}

	group, err := hub.LocalGroupCache.Get(GroupCacheKey(id))
	if err != nil {
		log.Errorf("获取群组数据失败，id:%d", id)
		return nil
	}
	return group.(*comp.Group)
}

// ExistsActivity 是否存在活动
func (c CacheSystem) ExistsActivity(id uint64) bool {
	return hub.LocalActivityCache.Has(ActivityCacheKey(id))
}

func (c CacheSystem) Activity(id uint64) *comp.Activity {
	if !c.ExistsActivity(id) {
		log.Infof("从 DB 获取活动数据，id:%d", id)
		act := mysql.Activity.GetActivityById(id)
		act.OutDB()
		err := hub.LocalActivityCache.Set(ActivityCacheKey(id), act)
		if err != nil {
			log.Errorf("从 DB 获取活动数据失败，id:%d", id)
			return nil
		}
	}

	act, err := hub.LocalActivityCache.Get(ActivityCacheKey(id))
	if err != nil {
		log.Errorf("获取活动数据失败，id:%d", id)
		return nil
	}
	return act.(*comp.Activity)
}
