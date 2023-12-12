package model

import (
	"gorm.io/gorm"
)

const TableNameMatch = "matches"

type Match struct {
	gorm.Model
	Name string `gorm:"column:name;comment:比赛名称" json:"name"` // 比赛名称
}

func (*Match) TableName() string {
	return TableNameMatch
}
