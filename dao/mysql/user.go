package mysql

import (
	"fmt"

	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v1/comp"
)

type UserDao struct {
}

var User = new(UserDao)

func (u UserDao) Create() {
	//fields := "username,password,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups"
	//sql := "INSERT INTO user (?) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	//result, err := db.Exec(sql, $fields, )
}

func (u *UserDao) GetUserByName(name string) {
	//
}

// GetIDs 获取所有记录的 ID
func (u *UserDao) GetIDs() []*comp.ID {
	var ids = make([]*comp.ID, 0)
	if err := db.Table(comp.TABLE_USER).Find(&ids).Error; err != nil {
		log.Errorf("获取用户数据出错，err: %s", err)
	}

	return ids
}

// GetUserByID 获取 ID 获取用户信息
func (u *UserDao) GetUserByID(id uint64) *comp.User {
	fields := "id,scores,username,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups,create_at"
	sql := fmt.Sprintf("SELECT %s FROM `user` WHERE id = ?", fields)

	user := comp.NewUser()
	if err := db.Raw(sql, id).Scan(user).Error; err != nil {
		log.Errorf("获取用户数据出错, ID:%d, err: %s", id, err)
		return nil
	}

	return user
}

func (u *UserDao) GetUsersByIds(ids string) []*comp.User {
	fields := "id,scores,username,token,nick,wx_token,wx_nick,sex,phone,email,ip,activities,groups,create_at"
	sql := fmt.Sprintf("SELECT %s FROM `user` WHERE id in(?)", fields)

	var users []*comp.User
	if err := db.Select(&users, sql, ids); err != nil {
		// gl.App.Logger().Errorf("Get user data failed, err: %v\n", err)
	}

	return users
}

func (u UserDao) UpdateUserById(id int64) {

}
