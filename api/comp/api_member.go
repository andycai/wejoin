package comp

import "time"

type APIMember struct {
	ID       int32     `json:"id"`
	UserID   int32     `json:"user_id"`
	GroupID  int32     `json:"group_id"`
	Scores   int32     `json:"scores"`
	Position int32     `json:"position"`
	Sex      int32     `json:"sex"`
	Avatar   string    `json:"avatar"`
	EnterAt  time.Time `json:"enter_at"`
	WxNick   string    `json:"wx_nick"`
	Nick     string    `json:"nick"`
}
