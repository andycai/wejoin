package main

import (
	"fmt"

	"github.com/andycai/axe-fiber/dao/mysql"
	"github.com/andycai/axe-fiber/v1/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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

	mysql.InitMySQL()

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n​",
		// Format:     "${pid} ${status} - ${method} ${path}\n",
		// TimeFormat: "02-Jan-2006",
		// TimeZone:   "America/New_York",
	}))

	router.InitRouter(app)
	app.Listen(viper.GetString("httpServer.addr"))
}
