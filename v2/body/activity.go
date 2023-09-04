package body

import (
	"time"

	"github.com/andycai/axe-fiber/v2/model"
)

type Activity struct {
	ID        int32  `json:"id" xml:"id" form:"id"`
	Kind      int32  `json:"kind" xml:"kind" form:"kind"`
	Type      int32  `json:"type" xml:"type" form:"type"`
	Title     string `json:"title" xml:"title" form:"title"`
	Remark    string `json:"remark" xml:"remark" form:"remark"`
	Quota     int32  `json:"quota" xml:"quota" form:"quota"`
	Waiting   int32  `json:"waiting" xml:"waiting" form:"waiting"`
	Status    int32  `json:"status" xml:"status" form:"status"`
	FeeType   int32  `json:"fee_type" xml:"fee_type" form:"fee_type"`
	FeeMale   int32  `json:"fee_male" xml:"fee_male" form:"fee_male"`
	FeeFemale int32  `json:"fee_female" xml:"fee_female" form:"fee_female"`
	Addr      string `json:"addr" xml:"addr" form:"addr"`
	Ahead     int32  `json:"ahead" xml:"ahead" form:"ahead"`
	BeginAt   string `json:"begin_at" xml:"begin_at" form:"begin_at"`
	EndAt     string `json:"end_at" xml:"end_at" form:"end_at"`
	Action    int32  `json:"action" xml:"action" form:"action"`
}

func (a Activity) ToModel() model.Activity {
	beginAt, _ := time.ParseInLocation("2006-01-02 15:04:05", a.BeginAt, time.Local)
	endAt, _ := time.ParseInLocation("2006-01-02 15:04:05", a.EndAt, time.Local)
	return model.Activity{
		ID:        a.ID,
		Kind:      a.Kind,
		Type:      a.Type,
		Title:     a.Title,
		Remark:    a.Remark,
		Quota:     a.Quota,
		Waiting:   a.Waiting,
		Status:    a.Status,
		FeeType:   a.FeeType,
		FeeMale:   a.FeeMale,
		FeeFemale: a.FeeFemale,
		Addr:      a.Addr,
		Ahead:     a.Ahead,
		BeginAt:   beginAt,
		EndAt:     endAt,
	}
}
