// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameActivity = "activity"

// Activity mapped from table <activity>
type Activity struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                    // 活动ID
	Planner   int32     `gorm:"column:planner;not null" json:"planner"`                               // 组织者ID
	GroupID   int32     `gorm:"column:group_id;not null" json:"group_id"`                             // 群组ID
	Kind      int32     `gorm:"column:kind;not null;default:1" json:"kind"`                           // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      int32     `gorm:"column:type;not null;default:1" json:"type"`                           // 活动类型:1全局保护,2全局公开,3群组
	Title     string    `gorm:"column:title;not null" json:"title"`                                   // 活动标题
	Remark    string    `gorm:"column:remark" json:"remark"`                                          // 活动描述
	Quota     int32     `gorm:"column:quota;not null;default:1" json:"quota"`                         // 名额
	Waiting   int32     `gorm:"column:waiting;not null;default:1" json:"waiting"`                     // 候补数量限制
	Status    int32     `gorm:"column:status;not null;default:1" json:"status"`                       // 活动状态:1进行中,2正常结算完成,3手动终止
	FeeType   int32     `gorm:"column:fee_type;not null;default:1" json:"fee_type"`                   // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   int32     `gorm:"column:fee_male;not null" json:"fee_male"`                             // 男费用,单位:分
	FeeFemale int32     `gorm:"column:fee_female;not null" json:"fee_female"`                         // 女费用,单位:分
	Addr      string    `gorm:"column:addr" json:"addr"`                                              // 活动地址
	Ahead     int32     `gorm:"column:ahead;not null" json:"ahead"`                                   // 可提前取消时间(小时)
	BeginAt   time.Time `gorm:"column:begin_at;not null" json:"begin_at"`                             // 开始时间
	EndAt     time.Time `gorm:"column:end_at;not null" json:"end_at"`                                 // 结束时间
	CreateAt  time.Time `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP" json:"create_at"` // 创建时间
	UpdateAt  time.Time `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP" json:"update_at"` // 更新时间
	DeleteAt  time.Time `gorm:"column:delete_at" json:"delete_at"`                                    // 删除时间
}

// TableName Activity's table name
func (*Activity) TableName() string {
	return TableNameActivity
}
