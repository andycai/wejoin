package dao

import "gorm.io/gorm"

var Q *Query

func InitQuery(db *gorm.DB) {
	Q = Use(db)
}
