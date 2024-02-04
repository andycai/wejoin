package model

import (
	"time"

	"gorm.io/gorm"
)

// User mapped from table <user>
type User struct {
	gorm.Model
	Username string    `gorm:"column:username;not null;comment:用户名" json:"username"`                 // 用户名
	Password string    `gorm:"column:password;not null;default:123456;comment:密码" json:"password"`   // 密码
	Scores   uint      `gorm:"column:scores;not null;comment:积分" json:"scores"`                      // 积分
	Token    string    `gorm:"column:token;not null;comment:token" json:"token"`                     // token
	WxToken  string    `gorm:"column:wx_token;not null;comment:微信session_key" json:"wx_token"`       // 微信session_key
	WxNick   string    `gorm:"column:wx_nick;not null;comment:微信昵称" json:"wx_nick"`                  // 微信昵称
	Nick     string    `gorm:"column:nick;not null;comment:昵称" json:"nick"`                          // 昵称
	Avatar   string    `gorm:"column:avatar;comment:头像" json:"avatar"`                               // 头像
	Gender   uint8     `gorm:"column:gender;not null;default:1;comment:性别:1男,2女,3自定义" json:"gender"` // 性别:1男,2女,3自定义
	Phone    string    `gorm:"column:phone;comment:手机号码" json:"phone"`                               // 手机号码
	Email    string    `gorm:"column:email;comment:邮箱" json:"email"`                                 // 邮箱
	Addr     string    `gorm:"column:addr;comment:住址" json:"addr"`                                   // 住址
	IP       string    `gorm:"column:ip;default:0;comment:ip地址" json:"ip"`                           // ip地址
	LoginAt  time.Time `gorm:"column:login_at;comment:登录时间" json:"login_at"`                         // 登录时间
	LogoutAt time.Time `gorm:"column:logout_at;comment:登出时间" json:"logout_at"`                       // 时间
}
