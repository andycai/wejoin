package body

type User struct {
	ID       int32  `json:"id" xml:"id" form:"id"`
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Token    string `json:"token" xml:"token" form:"token"`
	Nick     string `json:"nick" xml:"nick" form:"nick"`
	Avatar   string `json:"avatar" xml:"avatar" form:"avatar"`
	Gender   int32  `json:"gender" xml:"gender" form:"gender"`
	Phone    string `json:"phone" xml:"phone" form:"phone"`
	Email    string `json:"email" xml:"email" form:"email"`
	Addr     string `json:"addr" xml:"addr" form:"addr"`
	Action   int32  `json:"action" xml:"action" form:"action"`
}
