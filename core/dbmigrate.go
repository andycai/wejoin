package core

import (
	"github.com/andycai/wejoin/model"
	"gorm.io/gorm"
)

func AutoMigrate(dbs []*gorm.DB) {
	for _, db := range dbs {
		db.AutoMigrate(
			&model.User{},
			&model.Group{},
			&model.GroupApplication{},
			&model.GroupMember{},
			&model.Activity{},
			&model.ActivityUser{},
			&model.Match{},
		)
	}
}
