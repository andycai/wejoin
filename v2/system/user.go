package system

import (
	"errors"

	"github.com/andycai/axe-fiber/enum/action"
	"github.com/andycai/axe-fiber/v2/body"
	"github.com/andycai/axe-fiber/v2/model"
	"github.com/spf13/cast"

	"github.com/andycai/axe-fiber/v2/comp"
	"github.com/andycai/axe-fiber/v2/dao"
)

type UserSystem struct {
}

var User = new(UserSystem)

// GetInfo 获取用户信息
func (us UserSystem) GetInfo(uid int32) (*comp.APIUser, error) {
	u := dao.User
	info := &comp.APIUser{}
	err := u.Where(u.ID.Eq(uid)).Scan(info)
	if info.ID == 0 {
		err = errors.New("not found user data")
	}

	return info, err
}

// Register 注册
func (us UserSystem) Register(username, password, IP string) (error, int32) {
	u := dao.User
	err := u.Create(&model.User{
		Username: username,
		Password: password,
		Scores:   0,
		IP:       IP,
	})

	return err, cast.ToInt32(u.ID)
}

// Login 登录
func (us UserSystem) Login(username, password string) error {
	//
	return nil
}

// LoginJWT 登录
func (us UserSystem) LoginJWT(username, token string) error {
	return nil
}

// Update 更新用户信息
func (us UserSystem) Update(user *body.User) error {
	u := dao.User
	switch user.Action {
	case action.UserUpdatePassword:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Password, user.Password)
		return err
	case action.UserUpdateNick:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Nick, user.Nick)
		return err
	case action.UserUpdateAvatar:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Avatar, user.Avatar)
		return err
	case action.UserUpdateGender:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Gender, user.Gender)
		return err
	case action.UserUpdatePhone:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Phone, user.Phone)
		return err
	case action.UserUpdateEmail:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Email, user.Email)
		return err
	case action.UserUpdateAddr:
		_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Addr, user.Addr)
		return err
	}
	return nil
}

// Exists 用户是否存在
func (us UserSystem) Exists(uid int32) bool {
	u := dao.User
	count, err := u.Where(u.ID.Eq(uid)).Count()
	if err != nil {
		return false
	}

	return count > 0
}
