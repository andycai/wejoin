package body

import (
	"time"

	"github.com/andycai/axe-fiber/model"
)

type Activity struct {
	ID          uint   `json:"id" xml:"id" form:"id"`
	Kind        uint   `json:"kind" xml:"kind" form:"kind"`
	Type        uint   `json:"type" xml:"type" form:"type"`
	Title       string `json:"title" xml:"title" form:"title"`
	Description string `json:"description" xml:"description" form:"description"`
	Quota       uint   `json:"quota" xml:"quota" form:"quota"`
	Waiting     uint   `json:"waiting" xml:"waiting" form:"waiting"`
	Stage       uint   `json:"stage" xml:"stage" form:"stage"`
	FeeType     uint   `json:"fee_type" xml:"fee_type" form:"fee_type"`
	FeeMale     uint   `json:"fee_male" xml:"fee_male" form:"fee_male"`
	FeeFemale   uint   `json:"fee_female" xml:"fee_female" form:"fee_female"`
	Addr        string `json:"addr" xml:"addr" form:"addr"`
	Ahead       uint   `json:"ahead" xml:"ahead" form:"ahead"`
	BeginAt     string `json:"begin_at" xml:"begin_at" form:"begin_at"`
	EndAt       string `json:"end_at" xml:"end_at" form:"end_at"`
	Action      uint   `json:"action" xml:"action" form:"action"`
}

// ToModle 转换成 Model Activity
func (a Activity) ToModel() model.Activity {
	beginAt, _ := time.ParseInLocation("2006-01-02 15:04:05", a.BeginAt, time.Local)
	endAt, _ := time.ParseInLocation("2006-01-02 15:04:05", a.EndAt, time.Local)
	return model.Activity{
		ID:          a.ID,
		Kind:        a.Kind,
		Type:        a.Type,
		Title:       a.Title,
		Description: a.Description,
		Quota:       a.Quota,
		Waiting:     a.Waiting,
		Stage:       a.Stage,
		FeeType:     a.FeeType,
		FeeMale:     a.FeeMale,
		FeeFemale:   a.FeeFemale,
		Addr:        a.Addr,
		Ahead:       a.Ahead,
		BeginAt:     beginAt,
		EndAt:       endAt,
	}
}
