package model

import (
	"gorm.io/gorm"
)

type MatchMember struct {
	gorm.Model
	MatchID uint `gorm:"column:match_id;comment:比赛ID" json:"match_id"`
	UserID  uint `gorm:"column:user_id;comment:用户ID" json:"user_id"`
}
