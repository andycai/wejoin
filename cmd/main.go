package main

import (
	"path/filepath"

	"github.com/andycai/wejoin/api/middleware"
	"github.com/andycai/wejoin/conf"
	"github.com/andycai/wejoin/core"
	"github.com/andycai/wejoin/library/authentication"
	"github.com/andycai/wejoin/library/database"
	"github.com/andycai/wejoin/log"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	log.Setup()
	conf.ReadConf()

	// database open and init
	db, err := database.InitRDBMS(viper.GetString("db.type"),
		viper.GetString("db.dsn"),
		viper.GetInt("db.active"),
		viper.GetInt("db.idle"),
		viper.GetInt("db.idletimeout"))
	if err != nil {
		panic(err)
	}
	// dao.SetDefault(db)
	dbs := []*gorm.DB{db}
	core.AutoMigrate(dbs)
	core.SetupDatabase(dbs)
	authentication.SessionStart()
	core.SetZoneOffset(viper.GetInt("app.zoneoffset"))
	core.SetLang(viper.GetString("app.lang"))

	// Middleware
	middleware.Use(app)

	app.Static("/static", filepath.Join("", viper.GetString("app.static")))

	// router
	core.SetupRouter(app)

	err = app.Listen(viper.GetString("httpserver.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
