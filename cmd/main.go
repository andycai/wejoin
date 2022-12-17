package main

import (
	"github.com/andycai/axe-fiber/conf"
	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/log"
	"github.com/andycai/axe-fiber/v2/dao"
	"github.com/andycai/axe-fiber/v2/middleware"
	"github.com/andycai/axe-fiber/v2/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	log.Setup()
	conf.ReadConf() // 读取配置
	// mysql.ConnectMySQL() // 初始化数据库
	dao.InitQuery(mysql.CreateDB())
	middleware.Use(app) // 初始化中间件
	router.Setup(app)   // 初始化路由
	// system.Cache.InitIds() // 初始化数据表的 ID 到缓存

	err := app.Listen(viper.GetString("httpServer.addr"))
	if err != nil {
		panic(err)
	}
	defer func() {
		//
	}()
}
