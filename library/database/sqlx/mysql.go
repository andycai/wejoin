package sqlx

import (
	"fmt"
	"time"

	"github.com/andycai/axe-fiber/library/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQL(conf *database.Config) *sqlx.DB {
	db, err := sqlx.Open("mysql", conf.DSN)
	if err != nil {
		panic(fmt.Errorf("connect server failed, err: %v\n", err))
	}
	db.SetMaxOpenConns(conf.Active)
	db.SetMaxIdleConns(conf.Idle)
	db.SetConnMaxLifetime(conf.IdleTimeout * time.Second)

	return db
}
