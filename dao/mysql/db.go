package mysql

import (
	"github.com/andycai/axe-fiber/library/database"
	orm "github.com/andycai/axe-fiber/library/database/gorm"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func ConnectMySQL() {
	conf := &database.Config{
		DSN: viper.GetString("db.dsn"),
	}
	db = orm.NewMySQL(conf)
}
