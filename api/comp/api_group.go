package comp

import (
	"github.com/andycai/wejoin/constant"
	"github.com/andycai/wejoin/util/slice"
)

type APIGroup struct {
	ID         int32       `json:"id"`
	Level      int         `json:"level"`
	Scores     int         `json:"scores"`
	Name       string      `json:"name"`
	Logo       string      `json:"logo"`
	Notice     string      `json:"notice"`
	Addr       string      `json:"addr"`
	Activities []int32     `json:"activities"`
	Pending    []int32     `json:"pending"` // 申请入群列表
	Members    []APIMember `json:"members"`
}

// NewGroup 新建群组
func NewGroup() *APIGroup {
	g := new(APIGroup)
	return g
}

// OutDB 反序列化
func (g *APIGroup) OutDB() {
	// json.Unmarshal([]byte(g.Activities), &g.ActivitiesV)
	// json.Unmarshal([]byte(g.Pending), &g.PendingV)
	// json.Unmarshal([]byte(g.Members), &g.MembersV)
}

// NotInPendingV 不在申请列表
func (g APIGroup) NotInPendingV(index int) bool {
	return index < 0 || index >= len(g.Pending)
}

// IsMember 是否群成员
func (g APIGroup) IsMember(uid int32) bool {
	for _, member := range g.Members {
		if member.UserID == uid {
			return true
		}
	}
	return false
}

// IsOwner 是否群主
func (g APIGroup) IsOwner(uid int32) bool {
	for _, member := range g.Members {
		if member.UserID == uid && member.Position == constant.PositionGroupOwner {
			return true
		}
	}
	return false
}

// IsManager 是否管理员
func (g APIGroup) IsManager(uid int32) bool {
	for _, member := range g.Members {
		if member.UserID == uid && member.Position == constant.PositionGroupManager {
			return true
		}
	}
	return false
}

// ManagerCount 管理员数量
func (g APIGroup) ManagerCount() int {
	c := 0
	for _, member := range g.Members {
		if member.Position > constant.PositionGroupMember {
			c += 1
		}
	}
	return c
}

// ExistsActivity 是否存在活动
func (g *APIGroup) ExistsActivity(aid int32) bool {
	for _, v := range g.Activities {
		if v == aid {
			return true
		}
	}
	return false
}

// AddActivity 添加活动
func (g APIGroup) AddActivity(aid int32) {
	if !g.ExistsActivity(aid) {
		g.Activities = append(g.Activities, aid)
	}
}

// RemoveActivity 移除活动
func (g *APIGroup) RemoveActivity(aid int32) bool {
	if g.ExistsActivity(aid) {
		g.Activities = slice.Remove(g.Activities, aid)
		return true
	}
	return false
}

// Promote 提升管理员
func (g *APIGroup) Promote(uid int32) bool {
	for _, member := range g.Members {
		if member.UserID == uid {
			member.Position = constant.PositionGroupManager
			return true
		}
	}
	return false
}

// Transfer 转让群主
func (g *APIGroup) Transfer(uid, mid int32) bool {
	b := false
	if !g.IsMember(uid) || !g.IsMember(mid) {
		return false
	}
	for _, member := range g.Members {
		// 外部自行判断权限
		if member.UserID == uid {
			member.Position = constant.PositionGroupMember
		}
		if member.UserID == mid {
			member.Position = constant.PositionGroupOwner
			b = true
		}
	}
	return b
}

// Remove 移除群组成员
func (g *APIGroup) Remove(uid int32) bool {
	index := -1
	for i, member := range g.Members {
		if member.UserID == uid {
			index = i
			break
		}
	}
	if index >= 0 {
		g.Members = append(g.Members[:index], g.Members[index+1:]...)
		return true
	}
	return false
}

// NotIn 不在群组
func (g APIGroup) NotIn(uid int32) bool {
	for _, member := range g.Members {
		if member.UserID == uid {
			return false
		}
	}
	return true
}
