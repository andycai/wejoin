package model

import (
	"time"
)

// GroupMember mapped from table <group_member>
type GroupMember struct {
	ID       uint      `gorm:"column:id;primaryKey;autoIncrement:true;comment:成员id" json:"id"`  // 成员id
	GroupID  uint      `gorm:"column:group_id;not null;comment:群组id" json:"group_id"`           // 群组id
	UserID   uint      `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`             // 用户id
	Scores   uint      `gorm:"column:scores;not null;comment:积分" json:"scores"`                 // 积分
	Position uint      `gorm:"column:position;not null;default:1;comment:群组职位" json:"position"` // 群组职位
	Nick     string    `gorm:"column:nick;not null;comment:群组昵称" json:"nick"`                   // 群组昵称
	EnterAt  time.Time `gorm:"column:enter_at;comment:进入群组时间" json:"enter_at"`                  // 进入群组时间
	CreateAt time.Time `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`         // 创建时间
	UpdateAt time.Time `gorm:"column:update_at;not null;comment:更新时间" json:"update_at"`         // 更新时间
	DeleteAt time.Time `gorm:"column:delete_at;comment:删除时间" json:"delete_at"`                  // 删除时间
}
