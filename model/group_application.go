package model

import (
	"gorm.io/gorm"
)

// GroupApplication mapped from table <group_application>
type GroupApplication struct {
	gorm.Model
	UserID  uint `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`   // 用户id
	GroupID uint `gorm:"column:group_id;not null;comment:群组id" json:"group_id"` // 群组id
}
