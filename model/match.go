package model

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Name      string    `gorm:"column:name;comment:比赛名称" json:"name"`
	TypeID    uint      `gorm:"column:type_id;comment:比赛类型" json:"type_id"`
	Judge     uint      `gorm:"column:judge;comment:裁判" json:"judge"`
	MatchTime time.Time `gorm:"column:match_time;comment:比赛时间" json:"match_time"`
}
