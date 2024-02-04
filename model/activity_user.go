package model

import (
	"gorm.io/gorm"
)

// ActivityUser mapped from table <activity_user>
type ActivityUser struct {
	gorm.Model
	ActivityID uint   `gorm:"column:activity_id;not null;comment:活动id" json:"activity_id"`       // 活动id
	UserID     uint   `gorm:"column:user_id;not null;comment:报名用户id" json:"user_id"`             // 报名用户id
	Nick       string `gorm:"column:nick;comment:报名昵称" json:"nick"`                              // 报名昵称
	IsFriend   uint8  `gorm:"column:is_friend;not null;default:1;comment:是否朋友" json:"is_friend"` // 是否朋友
}
