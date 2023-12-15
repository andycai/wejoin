package activity

import (
	"github.com/andycai/axe-fiber/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "activity.gorm.db"
	KeyPostNoCheckRouter = "activity.router.nocheck"
	KeyPostCheckRouter   = "activity.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	api := r.Group("/v2/activities")
	{
		api.Get("/activities/:aid", GetActivityByID)
		api.Get("/activities", GetActivitiesByPage)

		api.Post("/activities", Create)
		api.Post("/activities/end", End)
		api.Post("/activities/apply", Apply)
		api.Post("/activities/cancel", Cancel)

		api.Put("/activities/:aid", Update)

		api.Delete("/activities/:aid", Remove)
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
