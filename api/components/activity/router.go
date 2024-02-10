package activity

import (
	"github.com/andycai/wejoin/core"
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
	api := r.Group("/v2")
	{
		api.Get("/activities/:id", handleGetByID)
		api.Get("/activities", handleGetByPage)
		api.Get("/activities/group/:gid", handleGetByGroupID)
		api.Get("/activities/user/:uid", handleGetByUserID)
		api.Get("/activities/organizer/:uid", handleGetByOrganizerUserID)

		api.Post("/activities", handleCreate)
		api.Post("/activities/complete", handleComplete)
		api.Post("/activities/end", handleEnd)
		api.Post("/activities/apply", handleApply)
		api.Post("/activities/cancel", handleCancel)

		api.Put("/activities/:aid", handleUpdate)

		api.Delete("/activities/:aid", handleRemove)
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
