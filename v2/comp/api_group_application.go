package comp

type APIGroupApplication struct {
	ID      int32 `json:"id"`       // 申请id
	UserID  int32 `json:"user_id"`  // 用户id
	GroupID int32 `json:"group_id"` // 群组id
}
