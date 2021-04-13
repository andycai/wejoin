package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func ConnectMySQL() (err error) {
	db, err = sqlx.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {
		panic(fmt.Errorf("connect server failed, err: %v\n", err))
	}
	db.SetMaxOpenConns(viper.GetInt("db.active"))
	db.SetMaxIdleConns(viper.GetInt("db.idle"))

	return
}

func Close() {
	db.Close()
}
