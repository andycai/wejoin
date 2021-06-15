package mysql

import (
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type ActivityDao struct {
	//
}

var Activity = new(ActivityDao)

func (a ActivityDao) Create() {
	//
}

// GetIDs 获取所有记录的 ID
func (a *ActivityDao) GetIDs() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_ACTIVITY).Find(&ids).Error; err != nil {
		log.Errorf("获取活动数据出错，err: %s", err)
	}

	return ids
}

func (a ActivityDao) GetActivityById(id uint64) *comp.Activity {
	act := comp.NewActivity()
	if err := db.Table(comp.TABLE_ACTIVITY).Where("id = ?", id).Find(act).Error; err != nil {
		log.Errorf("获取活动数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return act
}

func (a ActivityDao) GetActivitiesByType(typ, status, page, num int) {

}

func (a ActivityDao) GetActivitiesById(ids string) {

}

func (a ActivityDao) UpdateActivityStatus(id int64) {

}

func (a ActivityDao) UpdateActivityById(id int64) {

}
