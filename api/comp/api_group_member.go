package comp

import "time"

type APIGroupMember struct {
	ID       int32     `json:"id"`       // 成员id
	GroupID  int32     `json:"group_id"` // 群组id
	UserID   int32     `json:"user_id"`  // 用户id
	Scores   int32     `json:"scores"`   // 积分
	Position int32     `json:"position"` // 群组职位
	Alias_   string    `json:"alias"`    // 群组中的昵称
	Avatar   string    `json:"avatar"`   // 头像
	EnterAt  time.Time `json:"enter_at"` // 进入群组时间
}
