package model

import (
	"gorm.io/gorm"
)

// Group mapped from table <group>
type Group struct {
	gorm.Model
	Scores uint   `gorm:"column:scores;not null;comment:积分" json:"scores"`           // 积分
	Level  uint   `gorm:"column:level;not null;default:1;comment:群组等级" json:"level"` // 群组等级
	Name   string `gorm:"column:name;not null;comment:群组名称" json:"name"`             // 群组名称
	Logo   string `gorm:"column:logo;comment:群组LOGO" json:"logo"`                    // 群组LOGO
	Notice string `gorm:"column:notice;comment:群组公告" json:"notice"`                  // 群组公告
	Addr   string `gorm:"column:addr;comment:群组总部地址" json:"addr"`                    // 群组总部地址
}
