package model

import (
	"gorm.io/gorm"
)

const TableNameMatchMember = "match_member"

type MatchMember struct {
	gorm.Model
	MatchID uint `gorm:"column:match_id;comment:比赛ID" json:"match_id"`
	UserID  uint `gorm:"column:user_id;comment:用户ID" json:"user_id"`
}

func (*MatchMember) TableName() string {
	return TableNameMatchMember
}
