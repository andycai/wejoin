package user

import (
	"errors"

	"github.com/andycai/wejoin/api/body"
	"github.com/andycai/wejoin/model"
	"gorm.io/gorm"
)

type UserDao struct {
}

var Dao = new(UserDao)

// GetByID get the user by id
func (us UserDao) GetByID(uid uint) (*model.User, error) {
	userVo := model.User{}
	err := db.Raw(SqlQueryUserByID, uid).Scan(&userVo).Error

	return &userVo, err
}

// Register register
func (us UserDao) Register(username, password, ip string) (error, uint) {
	userVo := model.User{
		Username: username,
		Password: password,
		Scores:   0,
		IP:       ip,
	}
	err := db.Create(&userVo).Error

	return err, userVo.ID
}

// Login login
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
	// u := dao.User
	// switch user.Action {
	// case action.UserUpdatePassword:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Password, user.Password)
	// 	return err
	// case action.UserUpdateNick:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Nick, user.Nick)
	// 	return err
	// case action.UserUpdateAvatar:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Avatar, user.Avatar)
	// 	return err
	// case action.UserUpdateGender:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Gender, user.Gender)
	// 	return err
	// case action.UserUpdatePhone:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Phone, user.Phone)
	// 	return err
	// case action.UserUpdateEmail:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Email, user.Email)
	// 	return err
	// case action.UserUpdateAddr:
	// 	_, err := u.Where(u.ID.Eq(user.ID)).Update(u.Addr, user.Addr)
	// 	return err
	// }
	return nil
}

// ExistsByID 用户是否存在
func (us UserDao) ExistsByID(uid uint) error {
	err := db.Unscoped().Raw(SqlQueryUserByID, uid).First(&model.User{}).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
