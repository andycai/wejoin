package action

const (
	UserUpdatePassword int32 = iota + 100
	UserUpdateNick
	UserUpdateAvatar
	UserUpdateGender
	UserUpdatePhone
	UserUpdateEmail
	UserUpdateAddr
)

const (
	GroupUpdateName int32 = iota + 200
	GroupUpdateLogo
	GroupUpdateNotice
	GroupUpdateAddr
)
