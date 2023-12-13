package user

import (
	"github.com/andycai/axe-fiber/v2/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "user.gorm.db"
	KeyPostNoCheckRouter = "user.router.nocheck"
	KeyPostCheckRouter   = "user.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	api := r.Group("/v2/users")
	{
		api.Get("/:uid", GetUser)
		api.Put("/:uid", Update)
	}
}

func initCheckRouter(r fiber.Router) {
	//
}

func init() {
	core.RegisterDatabase(KeyPostDB, initDB)
	core.RegisterNoCheckRouter(KeyPostNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyPostCheckRouter, initCheckRouter)
}
