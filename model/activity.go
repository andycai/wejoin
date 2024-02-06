package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	UserID      uint      `gorm:"column:user_id;not null;comment:组织者ID" json:"user_id"`                                                      // 组织者ID
	GroupID     uint      `gorm:"column:group_id;not null;comment:群组ID" json:"group_id"`                                                     // 群组ID
	Kind        uint      `gorm:"column:kind;not null;default:1;comment:活动分类:1羽毛球,2篮球,3足球,4聚餐..." json:"kind"`                               // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type        uint      `gorm:"column:type;not null;default:1;comment:活动类型:1全局保护,2全局公开,3群组" json:"type"`                                   // 活动类型:1全局保护,2全局公开,3群组
	Title       string    `gorm:"column:title;not null;comment:活动标题" json:"title"`                                                           // 活动标题
	Description string    `gorm:"column:description;not null;comment:活动描述" json:"description"`                                               // 活动描述
	Quota       uint      `gorm:"column:quota;not null;default:1;comment:报名名额" json:"quota"`                                                 // 报名名额
	Waiting     uint      `gorm:"column:waiting;not null;default:1;comment:候补数量限制" json:"waiting"`                                           // 候补数量限制
	Stage       uint      `gorm:"column:stage;not null;default:1;comment:活动阶段:1报名阶段,2活动阶段,3正常完成和结算,4手动终止活动" json:"stage"`                    // 活动阶段:1报名阶段,2活动阶段,3正常完成和结算,4手动终止活动
	FeeType     uint      `gorm:"column:fee_type;not null;default:1;comment:结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定" json:"fee_type"` // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale     uint      `gorm:"column:fee_male;not null;comment:男费用,单位:分" json:"fee_male"`                                                 // 男费用,单位:分
	FeeFemale   uint      `gorm:"column:fee_female;not null;comment:女费用,单位:分" json:"fee_female"`                                             // 女费用,单位:分
	Address     string    `gorm:"column:address;comment:活动地址" json:"address"`                                                                // 活动地址
	Ahead       uint      `gorm:"column:ahead;not null;comment:可提前取消时间(小时)" json:"ahead"`                                                    // 可提前取消时间(小时)
	BeginAt     time.Time `gorm:"column:begin_at;not null;comment:开始时间" json:"begin_at"`                                                     // 开始时间
	EndAt       time.Time `gorm:"column:end_at;not null;comment:结束时间" json:"end_at"`                                                         // 结束时间
}
