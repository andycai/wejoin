package mysql

import (
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type GroupDao struct {
}

var Group = new(GroupDao)

// Create 创建群组
func (g *GroupDao) Create(group *comp.Group) uint64 {
	if result := db.Table(comp.TABLE_GROUP).Select("Level", "Name", "Members", "Activities", "Pending", "Notice", "Addr", "Logo").Create(group); result.Error != nil {
		log.Errorf("创建群组出错, name:%s, err: %s", group.Name, result.Error)
		return 0
	}

	return group.ID
}

// GetIds 获取所有记录的 ID
func (g *GroupDao) GetIds() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_GROUP).Find(&ids).Error; err != nil {
		log.Errorf("获取群组数据出错，err: %s", err)
	}

	return ids
}

// GetGroupById 根据 ID 获取群组信息
func (g GroupDao) GetGroupById(id uint64) *comp.Group {
	group := comp.NewGroup()
	if err := db.Table(comp.TABLE_GROUP).Where("id = ?", id).Find(group).Error; err != nil {
		log.Errorf("获取群组数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return group
}

// GetGroups 批量获取群组信息
func (g GroupDao) GetGroups(page, num int) []*comp.Group {
	var groups []*comp.Group
	if err := db.Table(comp.TABLE_GROUP).Limit(num).Offset(page * num).Or("id desc").Find(&groups).Error; err != nil {
		log.Errorf("获取群组数据出错, page:%d, num:%d, err: %s", page, num, err)
		return nil
	}

	return groups
}

// GetGroupsByIds 批量获取群组信息
func (g GroupDao) GetGroupsByIds(ids []uint64) []*comp.Group {
	var groups []*comp.Group
	if err := db.Table(comp.TABLE_GROUP).Where("id IN ?", ids).Find(&groups).Error; err != nil {
		log.Errorf("获取群组数据出错, ID:%v, err: %s", ids, err)
		return nil
	}

	return groups
}

// UpdateGroupById 更新群组信息
func (g GroupDao) UpdateGroupById(group *comp.Group) {
	if result := db.Table(comp.TABLE_USER).Select("Level", "Name", "Logo", "Notice", "Addr", "Members", "Pending", "Activities").Updates(group); result.Error != nil {
		log.Errorf("更新群组数据出错, ID:%d, err: %s", group.ID, result.Error)
	}
}
