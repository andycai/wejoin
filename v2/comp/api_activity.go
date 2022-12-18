package comp

import (
	"math"
	"time"

	"github.com/andycai/axe-fiber/enum"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
)

type APIActivity struct {
	ID        int32     `json:"id"`
	Planner   int32     `json:"planner"`
	Kind      int32     `json:"kind"`       // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      int32     `json:"type"`       // 活动类型:1全局保护,2全局公开,3群组
	Status    int32     `json:"status"`     // 活动状态:1进行中,2正常结算完成,3手动终止
	Quota     int32     `json:"quota"`      // 名额
	GroupId   int32     `json:"group_id"`   // 群组ID
	Ahead     int32     `json:"ahead"`      // 提前取消报名限制（小时）
	FeeType   int32     `json:"fee_type"`   // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   int32     `json:"fee_male"`   // 男费用
	FeeFemale int32     `json:"fee_female"` // 女费用
	Title     string    `json:"title"`
	Remark    string    `json:"remark"`
	Addr      string    `json:"addr"`
	BeginAt   time.Time `json:"begin_at"`
	EndAt     time.Time `json:"end_at"`
	Queue     []int32   `json:"queue"` // 报名队列
}

// IsPlanner 是否发起者
func (a APIActivity) IsPlanner(uid int32) bool {
	return uid == a.Planner
}

// Settle 结算
func (a *APIActivity) Settle(fee int32) {
	switch a.FeeType {
	case enum.FeeTypeActivityAA:
		cost := math.Round(cast.ToFloat64(fee) / cast.ToFloat64(len(a.Queue)))
		a.FeeMale = cast.ToInt32(cost)
		a.FeeFemale = cast.ToInt32(cost)
	case enum.FeeTypeActivityAB:
		// a.FeeFemale = cast.ToInt32(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeMale*a.maleCount())))
	case enum.FeeTypeActivityBA:
		// a.FeeMale = cast.ToInt32(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeFemale*a.femaleCount())))
	}
}

// OverQuota 报名的人数超过候补的限制，避免乱报名，如带100000人报名
func (a APIActivity) OverQuota(total int32) bool {
	return (cast.ToInt32((a.Queue)) + total - a.Quota) > enum.ActivityOverFlow
}

// NotEnough 要取消报名的数量超过已经报名的数量
func (a APIActivity) NotEnough(uid int32, total int) bool {
	c := 0
	for _, v := range a.Queue {
		if v == uid {
			c += 1
		}
	}
	return total > c
}

// InQueueV 在报名中
func (a APIActivity) InQueueV(uid int32) bool {
	for _, v := range a.Queue {
		if v == uid {
			return true
		}
	}
	return false
}

// GetIdFromQueueV 根据索引从报名队列中获取ID
func (a APIActivity) GetIdFromQueueV(index int) int32 {
	if index < 0 || index >= len(a.Queue) {
		return 0
	}
	return a.Queue[index]
}

// HasBegun 是否开始
func (a APIActivity) HasBegun() bool {
	return carbon.Parse(a.BeginAt.Format("2006-01-02 15:04:05")).DiffInHours() <= 0
}

// CanCancel 能否取消报名
func (a APIActivity) CanCancel() bool {
	hours := carbon.Parse(a.BeginAt.Format("2006-01-02 15:04:05")).DiffInHours()
	return cast.ToInt32(hours) >= a.Ahead
}

// 私有方法
