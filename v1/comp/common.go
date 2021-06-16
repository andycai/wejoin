package comp

type BodyObject struct {
	Token string `json:"token" form:"token"`

	// 用户
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Nick     string `json:"nick" form:"nick"`
	Sex      uint   `json:"sex" form:"sex"`

	// 分页
	Page uint `json:"page" form:"page"`
	Num  uint `json:"num" form:"num"`

	// 活动
	GroupId     uint64 `json:"group_id" form:"group_id"`
	Kind        uint   `json:"kind" form:"kind"`
	Type        uint   `json:"type" form:"type"`
	Status      uint   `json:"status" form:"status"`
	Quota       uint   `json:"quota" form:"quota"`
	Title       string `json:"title" form:"title"`
	Remark      string `json:"remark" form:"remark"`
	Fee         uint   `json:"fee" form:"fee"`
	FeeType     uint   `json:"fee_type" form:"fee_type"`
	FeeMale     uint   `json:"fee_male" form:"fee_male"`
	FeeFemale   uint   `json:"fee_female" form:"fee_female"`
	Ahead       uint   `json:"ahead" form:"ahead"`
	BeginAt     string `json:"begin_at" form:"begin_at"`
	EndAt       string `json:"end_at" form:"end_at"`
	MaleCount   uint   `json:"male_count" form:"male_count"`
	FemaleCount uint   `json:"female_count" form:"female_count"`

	// 群组
	GroupName string `json:"name" form:"name"`
	Logo      string `json:"logo" form:"logo"`
	Addr      string `json:"addr" form:"addr"`
	Notice    string `json:"notice" form:"notice"`
	Pass      bool   `json:"pass" form:"pass"`
	Index     uint   `json:"index" form:"index"`
}
