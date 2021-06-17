package mysql

import (
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type ActivityDao struct {
	//
}

var Activity = new(ActivityDao)

// Create 创建活动
func (a ActivityDao) Create(act *comp.Activity) uint64 {
	if result := db.Table(comp.TABLE_ACTIVITY).Select("Planner", "GroupId", "Kind", "Type", "Quota", "Title", "Remark", "Status", "FeeType", "FeeMale", "FeeFemale", "Queue", "QueueSex", "Addr", "Ahead", "BeginAt", "EndAt").Create(act); result.Error != nil {
		log.Errorf("创建活动出错, title:%s, err:%s", act.Title, result.Error)
		return 0
	}

	return act.ID
}

// GetIds 获取所有记录的 ID
func (a *ActivityDao) GetIds() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_ACTIVITY).Find(&ids).Error; err != nil {
		log.Errorf("获取活动数据出错，err: %s", err)
	}

	return ids
}

// GetActivityById 根据 ID 获取活动信息
func (a ActivityDao) GetActivityById(id uint64) *comp.Activity {
	act := comp.NewActivity()
	if err := db.Table(comp.TABLE_ACTIVITY).Where("id = ?", id).Find(act).Error; err != nil {
		log.Errorf("获取活动数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return act
}

// GetActivitiesByType 根据类型获取活动数据
func (a ActivityDao) GetActivitiesByType(typ, status, page, num int) []*comp.Activity {
	var acts []*comp.Activity
	if err := db.Table(comp.TABLE_ACTIVITY).Where("type = ? AND status = ?", typ, status).Limit(num).Offset(page * num).Order("id desc").Find(&acts).Error; err != nil {
		log.Errorf("获取活动数据出错, type:%d, err: %s", typ, err)
		return nil
	}

	return acts
}

// GetActivitiesByIds 批量获取活动数据
func (a ActivityDao) GetActivitiesByIds(ids []uint64) []*comp.Activity {
	var acts []*comp.Activity
	if err := db.Table(comp.TABLE_ACTIVITY).Where("id IN ?", ids).Find(&acts).Error; err != nil {
		log.Errorf("获取活动数据出错, ID:%v, err: %s", ids, err)
		return nil
	}

	return acts
}

// UpdateActivityStatus 更新活动状态
func (a ActivityDao) UpdateActivityStatus(act *comp.Activity) {
	if result := db.Table(comp.TABLE_GROUP).Select("Status", "FeeMale", "FeeFemale").Updates(act); result.Error != nil {
		log.Errorf("更新活动数据出错, ID:%d, err: %s", act.ID, result.Error)
	}
}

// UpdateActivityById 根据 ID 更新活动信息
func (a ActivityDao) UpdateActivityById(act *comp.Activity) {
	if result := db.Table(comp.TABLE_ACTIVITY).Select("Quota", "Title", "Remark", "Status", "Ahead", "Queue", "QueueSex", "FeeMale", "FeeFemale", "Addr", "BeginAt", "EndAt").Updates(act); result.Error != nil {
		log.Errorf("更新活动数据出错, ID:%d, err: %s", act.ID, result.Error)
	}
}
