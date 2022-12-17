package system

import (
	"fmt"

	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/hub"
)

// activityCacheKey 活动缓存 key
func activityCacheKey(id uint64) string {
	return fmt.Sprintf("Activity:%d", id)
}

// InitActivityIds 缓存活动的 id
func (c CacheSystem) InitActivityIds() {
	activityIDs := mysql.Activity.GetIds()
	for _, v := range activityIDs {
		hub.ActivityIDs[v.ID] = struct{}{}
	}
}

// ExistsActivity 是否存在活动
func (c CacheSystem) ExistsActivity(id uint64) bool {
	return hub.LocalActivityCache.Has(activityCacheKey(id))
}

func (c CacheSystem) Activity(id uint64) *comp.Activity {
	if !c.ExistsActivity(id) {
		log.Infof("从 DB 获取活动数据，id:%d", id)
		act := mysql.Activity.GetActivityById(id)
		act.OutDB()
		err := hub.LocalActivityCache.Set(activityCacheKey(id), act)
		if err != nil {
			log.Errorf("从 DB 获取活动数据失败，id:%d", id)
			return nil
		}
	}

	act, err := hub.LocalActivityCache.Get(activityCacheKey(id))
	if err != nil {
		log.Errorf("获取活动数据失败，id:%d", id)
		return nil
	}
	return act.(*comp.Activity)
}
