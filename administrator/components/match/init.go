package match

import (
	"github.com/andycai/wejoin/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyDB            = "admin.match.gorm.db"
	KeyNoCheckRouter = "admin.match.router.nocheck"
	KeyCheckRouter   = "admin.match.router.check"
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
