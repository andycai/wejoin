package gorm

import (
	"fmt"

	"github.com/andycai/axe-fiber/library/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(conf *database.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect server failed, err: %v\n", err))
	}
	return db
}
