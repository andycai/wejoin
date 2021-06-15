package mysql

import (
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type GroupDao struct {
}

var Group = new(GroupDao)

func (g *GroupDao) Create() {
	//
}

// GetIDs 获取所有记录的 ID
func (g *GroupDao) GetIDs() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_GROUP).Find(&ids).Error; err != nil {
		log.Errorf("获取群组数据出错，err: %s", err)
	}

	return ids
}

func (g GroupDao) GetGroupById(id uint64) *comp.Group {
	group := comp.NewGroup()
	if err := db.Table(comp.TABLE_GROUP).Where("id = ?", id).Find(group).Error; err != nil {
		log.Errorf("获取群组数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return group
}

func (g GroupDao) GetGroups(page, num int) {

}

func (g GroupDao) GetGroupsByIds(ids string) {

}

func (g GroupDao) UpdateGroupById(id int) {

}
