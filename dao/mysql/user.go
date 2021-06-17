package mysql

import (
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type UserDao struct {
}

var User = new(UserDao)

// Create 创建用户
func (u UserDao) Create(user *comp.User) uint64 {
	if result := db.Table(comp.TABLE_USER).Select("Userame", "Password", "Token", "Nick", "WxToken", "WxNick", "Sex", "Phone", "Email", "Ip", "Activities", "Groups").Create(user); result.Error != nil {
		log.Errorf("创建用户出错, name:%s, err: %s", user.Username, result.Error)
		return 0
	}

	return user.ID
}

// GetUserByName 根据名字获取用户信息
func (u *UserDao) GetUserByName(name string) *comp.User {
	user := comp.NewUser()
	if err := db.Table(comp.TABLE_USER).Where("username = ?", name).Find(user).Error; err != nil {
		log.Errorf("获取用户数据出错, name:%s, err: %s", name, err)
		return nil
	}

	return user
}

// GetIds 获取所有记录的 ID
func (u *UserDao) GetIds() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_USER).Find(&ids).Error; err != nil {
		log.Errorf("获取用户数据出错，err: %s", err)
	}

	return ids
}

// GetUserById 根据 ID 获取用户信息
func (u *UserDao) GetUserById(id uint64) *comp.User {
	user := comp.NewUser()
	if err := db.Table(comp.TABLE_USER).Where("id = ?", id).Find(user).Error; err != nil {
		log.Errorf("获取用户数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return user
}

// GetUsersByIds 批量获取用户信息
func (u *UserDao) GetUsersByIds(ids []uint64) []*comp.User {
	var users []*comp.User
	if err := db.Table(comp.TABLE_USER).Where("id IN ?", ids).Find(&users).Error; err != nil {
		log.Errorf("获取用户数据出错, ID:%v, err: %s", ids, err)
		return nil
	}

	return users
}

// UpdateUserById 更新用户信息
func (u UserDao) UpdateUserById(user *comp.User) {
	if result := db.Table(comp.TABLE_USER).Select("Token", "Nick", "WxToken", "WxNick", "Ip", "Activities", "Groups").Updates(user); result.Error != nil {
		log.Errorf("更新用户数据出错, ID:%d, err: %s", user.ID, result.Error)
	}
}
