package comp

import (
	"math"

	"github.com/andycai/axe-fiber/enum"
	"github.com/golang-module/carbon"
	"github.com/spf13/cast"
)

type Activity struct {
	ID        uint64   `gorm:"primaryKey" json:"id"`
	Planner   uint64   `json:"planner"`
	Kind      int      `json:"kind"`       // 活动分类:1羽毛球,2篮球,3足球,4聚餐...
	Type      int      `json:"type"`       // 活动类型:1全局保护,2全局公开,3群组
	Status    int      `json:"status"`     // 活动状态:1进行中,2正常结算完成,3手动终止
	Quota     int      `json:"quota"`      // 名额
	GroupId   uint64   `json:"group_id"`   // 群组ID
	Ahead     int      `json:"ahead"`      // 提前取消报名限制（小时）
	FeeType   int      `json:"fee_type"`   // 结算方式:1免费,2活动前,3活动后男女平均,4活动后男固定|女平摊,5活动后男平摊|女固定
	FeeMale   int      `json:"fee_male"`   // 男费用
	FeeFemale int      `json:"fee_female"` // 女费用
	Title     string   `json:"title"`
	Remark    string   `json:"remark"`
	Addr      string   `json:"addr"`
	BeginAt   string   `json:"begin_at"`
	EndAt     string   `json:"end_at"`
	Queue     string   `json:"-"`
	QueueSex  string   `json:"-"`
	QueueV    []uint64 `gorm:"-" json:"queue"`     // 报名队列
	QueueSexV []int    `gorm:"-" json:"queue_sex"` // 报名队列中的性别
}

// NewActivity 新建活动
func NewActivity() *Activity {
	a := new(Activity)
	return a
}

// OutDB 反序列化
func (a *Activity) OutDB() {
	json.Unmarshal([]byte(a.Queue), &a.QueueV)
	json.Unmarshal([]byte(a.QueueSex), &a.QueueSexV)
}

// InGroup 是否群组活动
func (a Activity) InGroup() bool {
	return a.GroupId > 0
}

// IsPlanner 是否发起者
func (a Activity) IsPlanner(uid uint64) bool {
	return uid == a.Planner
}

// Settle 结算
func (a *Activity) Settle(fee int) {
	switch a.FeeType {
	case enum.FeeTypeActivityAA:
		cost := math.Round(cast.ToFloat64(fee) / cast.ToFloat64(a.totalCount()))
		a.FeeMale = cast.ToInt(cost)
		a.FeeFemale = cast.ToInt(cost)
	case enum.FeeTypeActivityAB:
		a.FeeFemale = cast.ToInt(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeMale*a.maleCount())))
	case enum.FeeTypeActivityBA:
		a.FeeMale = cast.ToInt(math.Round(cast.ToFloat64(fee) - cast.ToFloat64(a.FeeFemale*a.femaleCount())))
	}
}

// OverQuota 报名的人数超过候补的限制，避免乱报名，如带100000人报名
func (a Activity) OverQuota(total int) bool {
	return len(a.QueueV)+total-a.Quota > enum.ActivityOverFlow
}

// NotEnough 要取消报名的数量超过已经报名的数量
func (a Activity) NotEnough(uid uint64, total int) bool {
	c := 0
	for _, v := range a.QueueV {
		if v == uid {
			c += 1
		}
	}
	return total > c
}

// InQueueV 在报名中
func (a Activity) InQueueV(uid uint64) bool {
	for _, v := range a.QueueV {
		if v == uid {
			return true
		}
	}
	return false
}

// GetIdFromQueueV 根据索引从报名队列中获取ID
func (a Activity) GetIdFromQueueV(index int) uint64 {
	if index < 0 || index >= len(a.QueueV) {
		return 0
	}
	return a.QueueV[index]
}

// Enqueue 报名入队
func (a *Activity) Enqueue(uid uint64, maleCount, femaleCount int) {
	a.fixQueueV()
	for i := 0; i < maleCount; i++ {
		a.QueueV = append(a.QueueV, uid)
		a.QueueSexV = append(a.QueueSexV, enum.TypeSexMale)
	}
	for i := 0; i < femaleCount; i++ {
		a.QueueV = append(a.QueueV, uid)
		a.QueueSexV = append(a.QueueSexV, enum.TypeSexFemale)
	}
}

// Dequeue 取消报名
func (a *Activity) Dequeue(index int) bool {
	a.fixQueueV()
	if index < 0 || index >= len(a.QueueV) {
		return false
	}
	a.QueueV = append(a.QueueV[:index], a.QueueV[index+1:]...)
	a.QueueSexV = append(a.QueueSexV[:index], a.QueueSexV[index+1:]...)
	return true
}

// DequeueMany 取消多个报名
func (a *Activity) DequeueMany(uid uint64, maleCount, femaleCount int) {
	a.fixQueueV()
	mCount := 0
	fCount := 0
	size := len(a.QueueV)
	posArr := make([]int, 1)
	for i := size - 1; i >= 0; i-- {
		if a.QueueV[i] == uid {
			// 男
			if a.QueueSexV[i] == enum.TypeSexMale && maleCount > mCount {
				mCount += 1
				posArr = append(posArr, i)
			}
			// 女
			if a.QueueSexV[i] == enum.TypeSexFemale && femaleCount > fCount {
				fCount += 1
				posArr = append(posArr, i)
			}
			if mCount >= maleCount && fCount >= femaleCount {
				break
			}
		}
	}
	for _, v := range posArr {
		a.QueueV = append(a.QueueV[:v], a.QueueV[v+1:]...)
		a.QueueSexV = append(a.QueueSexV[:v], a.QueueSexV[v+1:]...)
	}
}

// HasBegun 是否开始
func (a Activity) HasBegun() bool {
	return carbon.Parse(a.BeginAt).DiffInHours() <= 0
}

// CanCancel 能否取消报名
func (a Activity) CanCancel() bool {
	hours := carbon.Parse(a.BeginAt).DiffInHours()
	return cast.ToInt(hours) >= a.Ahead
}

// 私有方法

// 长度不一致，修正使其一致
func (a *Activity) fixQueueV() {
	df := len(a.QueueSexV) - len(a.QueueV)
	switch {
	case df > 0:
		a.QueueSexV = a.QueueSexV[:len(a.QueueV)]
	case df < 0:
		a.QueueV = a.QueueV[:len(a.QueueSexV)]
	}
}

// totalCount 有效的报名人数（不包括候选）
func (a Activity) totalCount() int {
	c := 0
	size := len(a.QueueV)
	if a.Quota >= size {
		c = size
	} else {
		c = a.Quota
	}
	return c
}

// maleCount 男性数量
func (a Activity) maleCount() int {
	c := 0
	total := a.totalCount()
	for i := 0; i < total; i++ {
		if len(a.QueueSexV) > i && a.QueueSexV[i] == enum.TypeSexMale {
			c += 1
		}
	}
	return c
}

// femaleCount 女性数量
func (a Activity) femaleCount() int {
	c := 0
	total := a.totalCount()
	for i := 0; i < total; i++ {
		if len(a.QueueSexV) > i && a.QueueSexV[i] == enum.TypeSexFemale {
			c += 1
		}
	}
	return c
}
