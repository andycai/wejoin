package match

import (
	"github.com/andycai/axe-fiber/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "match.gorm.db"
	KeyPostNoCheckRouter = "match.router.nocheck"
	KeyPostCheckRouter   = "match.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	api := r.Group("/v2")
	{
		api.Get("/matches/:gid", GetMatchByID)
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
