package group

import (
	"github.com/andycai/wejoin/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.group.gorm.db"
	KeyNoCheckRouter = "admin.group.router.nocheck"
	KeyCheckRouter   = "admin.group.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	admin := r.Group("/admin")
	{
		admin.Get("/login", LoginPage)
	}
}

func initCheckRouter(r fiber.Router) {
	r.Get("/logout", LogoutAction)
}

func init() {
	core.RegisterDatabase(KeyDB, initDB)
	core.RegisterNoCheckRouter(KeyNoCheckRouter, initNoCheckRouter)
	core.RegisterAdminCheckRouter(KeyCheckRouter, initCheckRouter)
}
