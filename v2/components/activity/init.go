package activity

import (
	"github.com/andycai/axe-fiber/v2/core"
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
		// 获取活动详细内容
		api.Get("/:aid", GetActivityByID)
		// 获取最近一段的活动
		api.Get("/list/:page/:num", GetActivities)
		// TODO: 获取最近一段时间附近的活动

		// 创建活动
		api.Post("/create", Create)
		// 手动终止活动
		api.Post("/end/:aid", End)
		// 申请参加活动
		api.Post("/apply/:aid", Apply)
		// 取消参加活动
		api.Post("/cancel/:aid", Cancel)

		// 修改活动内容
		api.Put("/:aid}", Update)

		// 删除活动
		api.Delete("/remove/:aid", Remove)
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
