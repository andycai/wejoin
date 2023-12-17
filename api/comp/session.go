package comp

import "time"

type Session struct {
	Uid     int32     `json:"uid"`
	Sex     int32     `json:"sex"`
	Token   string    `json:"token"`
	LoginAt time.Time `json:"login_at"`
}
