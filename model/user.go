package model

import "time"

type User struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Sex        int       `json:"sex"`
	Scores     int       `json:"scores"`
	Username   string    `json:"username"`
	Token      string    `json:"token"`
	WxToken    string    `json:"wx_token"`
	WxNick     string    `json:"wx_nick"`
	Nick       string    `json:"nick"`
	Ip         string    `json:"ip"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Groups     string    `json:"groups"`
	Activities string    `json:"activities"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	DeleteAt   time.Time `json:"delete_at"`
}
