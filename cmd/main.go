package main

import (
	"log"

	"github.com/andycai/axe-fiber/conf"
	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/gl"
	"github.com/andycai/axe-fiber/v1/middleware"
	"github.com/andycai/axe-fiber/v1/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()

	gl.Logger()
	conf.ReadConf()                 // 读取配置
	mysql.InitMySQL()               // 初始化数据库
	middleware.FiberMiddleware(app) // 初始化中间件
	router.InitRouter(app)          // 初始化路由

	log.Fatal(app.Listen(viper.GetString("httpServer.addr")))
}
