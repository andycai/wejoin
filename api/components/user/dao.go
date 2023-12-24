package user

import (
	"github.com/andycai/axe-fiber/api/body"
	"github.com/andycai/axe-fiber/api/dao"
	"github.com/andycai/axe-fiber/enum/action"
	"github.com/andycai/axe-fiber/model"
	"github.com/spf13/cast"
)

type UserDao struct {
}

var Dao = new(UserDao)

// GetInfo 获取用户信息
func (us UserDao) GetByID(uid uint) (*model.User, error) {
	userVo := model.User{}
	err := db.Raw(SqlQueryUserByID, uid).Scan(&userVo).Error

	return &userVo, err
}

// Register 注册
func (us UserDao) Register(username, password, ip string) (error, uint) {
	u := dao.User
	err := u.Create(&model.User{
		Username: username,
		Password: password,
		Scores:   0,
		IP:       ip,
	})

	return err, cast.ToInt32(u.ID)
}

// Login 登录
func (us UserDao) Login(username, password string) error {
	//
	return nil
}

// LoginJWT 登录
func (us UserDao) LoginJWT(username, token string) error {
	return nil
}

// Update 更新用户信息
func (us UserDao) Update(user *body.User) error {
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
func (us UserDao) Exists(uid uint) bool {
	u := dao.User
	count, err := u.Where(u.ID.Eq(uid)).Count()
	if err != nil {
		return false
	}

	return count > 0
}
