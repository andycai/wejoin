package system

import (
	"fmt"

	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
	"github.com/andycai/axe-fiber/v1/hub"
)

// groupCacheKey 群组缓存 key
func groupCacheKey(id uint64) string {
	return fmt.Sprintf("Group:%d", id)
}

// InitGroupIds 缓存群组的 id
func (c CacheSystem) InitGroupIds() {
	groupIDs := mysql.Group.GetIds()
	for _, v := range groupIDs {
		hub.GroupIDs[v.ID] = struct{}{}
	}
}

// ExistsGroup 是否存在群组
func (c CacheSystem) ExistsGroup(id uint64) bool {
	return hub.LocalGroupCache.Has(groupCacheKey(id))
}

func (c CacheSystem) Group(id uint64) *comp.Group {
	if !c.ExistsGroup(id) {
		log.Infof("从 DB 获取群组数据，id:%d", id)
		group := mysql.Group.GetGroupById(id)
		group.OutDB()
		err := hub.LocalGroupCache.Set(groupCacheKey(id), group)
		if err != nil {
			log.Errorf("从 DB 获取群组数据失败，id:%d", id)
			return nil
		}
	}

	group, err := hub.LocalGroupCache.Get(groupCacheKey(id))
	if err != nil {
		log.Errorf("获取群组数据失败，id:%d", id)
		return nil
	}
	return group.(*comp.Group)
}
