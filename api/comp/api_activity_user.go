package comp

type APIActivityUser struct {
	ID         int32  `json:"id"`          // 活动报名id
	ActivityID int32  `json:"activity_id"` // 活动id
	UserID     int32  `json:"user_id"`     // 报名用户id
	Alias_     string `json:"alias"`       // 报名昵称
	IsFriend   int32  `json:"is_friend"`   // 是否朋友
}
