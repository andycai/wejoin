package main

import (
	"axe/dao/mysql"
	"axe/gl"
	"axe/v1/router"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	pflag.String("app.cacheDir", "./cache/", "cache directory")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	//fmt.Println(viper.GetString("app.cacheDir"))

	mysql.InitMySQL()

	app := fiber.New()
	gl.App = app

	router.InitRouter(app)
	app.Listen(viper.GetString("httpServer.addr"))
}
