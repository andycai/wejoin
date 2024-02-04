package model

import (
	"gorm.io/gorm"
)

type MatchType struct {
	gorm.Model
	Name        string `gorm:"column:name;comment:比赛类型" json:"name"`
	Description string `gorm:"column:description;comment:描述" json:"description"`
}
