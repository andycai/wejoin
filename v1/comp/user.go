package comp

import (
	"time"

	"github.com/andycai/axe-fiber/util/slice"
)

type User struct {
	ID          uint64   `gorm:"primaryKey" json:"id"`
	Sex         int      `json:"sex"`
	Scores      int      `json:"scores"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Token       string   `json:"token"`
	WxToken     string   `json:"wx_token"`
	WxNick      string   `json:"wx_nick"`
	Nick        string   `json:"nick"`
	Ip          string   `json:"ip"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	CreateAt    string   `json:"create_at"`
	Groups      string   `json:"-"`
	Activities  string   `json:"-"`
	GroupsV     []uint64 `gorm:"-" json:"groups"`
	ActivitiesV []uint64 `gorm:"-" json:"activities"`
}

type MapUserIds = map[int64]struct{}

// NewUser 创建用户
func NewUser() *User {
	u := new(User)
	return u
}

// OutDB 出库后的数据格式
func (u *User) OutDB() {
	json.Unmarshal([]byte(u.Groups), &u.GroupsV)
	json.Unmarshal([]byte(u.Activities), &u.ActivitiesV)
	t, err := time.Parse(TimeFormatE, u.CreateAt)
	if err == nil {
		u.CreateAt = t.Format(TimeFormat)
	}
}

// InDB 入库前的数据格式
func (u User) InDB() {
	data, err := json.Marshal(u.GroupsV)
	if err == nil {
		u.Groups = string(data)
	}
	data, err = json.Marshal(u.ActivitiesV)
	if err == nil {
		u.Activities = string(data)
	}
}

// ExistsActivity 存在活动
func (u User) ExistsActivity(aid uint64) bool {
	for _, activity := range u.ActivitiesV {
		if activity == aid {
			return true
		}
	}
	return false
}

// AddActivity 添加活动
func (u *User) AddActivity(aid uint64) bool {
	if !u.ExistsActivity(aid) {
		u.ActivitiesV = append(u.ActivitiesV, aid)
		return true
	}
	return false
}

// RemoveActivity 移除活动
func (u *User) RemoveActivity(aid uint64) bool {
	if u.ExistsActivity(aid) {
		u.ActivitiesV = slice.RemoveU64(u.ActivitiesV, aid)
		return true
	}
	return false
}

// ExistsGroup 是否在群组
func (u User) ExistsGroup(gid uint64) bool {
	for _, group := range u.GroupsV {
		if group == gid {
			return true
		}
	}
	return false
}

// AddGroup 添加群组
func (u *User) AddGroup(gid uint64) bool {
	if !u.ExistsGroup(gid) {
		u.GroupsV = append(u.GroupsV, gid)
		return true
	}
	return false
}

// RemoveGroup 移除群组
func (u *User) RemoveGroup(gid uint64) bool {
	if u.ExistsGroup(gid) {
		u.GroupsV = slice.RemoveU64(u.GroupsV, gid)
		return true
	}
	return false
}
