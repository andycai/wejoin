package comp

type APIUser struct {
	ID         int32   `json:"id"`
	Sex        int     `json:"sex"`
	Scores     int     `json:"scores"`
	Username   string  `json:"username"`
	Token      string  `json:"token"`
	WxToken    string  `json:"wx_token"`
	WxNick     string  `json:"wx_nick"`
	Nick       string  `json:"nick"`
	Ip         string  `json:"ip"`
	Phone      string  `json:"phone"`
	Email      string  `json:"email"`
	Groups     []int32 `json:"groups"`
	Activities []int32 `json:"activities"`
}
