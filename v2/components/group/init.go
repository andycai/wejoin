package group

import (
	"github.com/andycai/axe-fiber/v2/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "group.gorm.db"
	KeyPostNoCheckRouter = "group.router.nocheck"
	KeyPostCheckRouter   = "group.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

func initNoCheckRouter(r fiber.Router) {
	api := r.Group("/v2/groups")
	{
		// 获取群组详细信息
		api.Get("/:gid", GetGroupByID)
		// 获取最近创建的群组
		api.Get("/list/:page/:num", GetGroups)
		// 获取申请加入群组的列表
		api.Get("/applications/:gid", GetApplyList)
		// 获取群组活动列表
		api.Get("/activities/:gid/:page/:num", GetActivitiesByGroupId)

		// 创建群组
		api.Post("/create", Create)
		// 申请加入群组
		api.Post("/apply/:gid/:uid", Apply)
		// 同意加入
		api.Post("/approve/:gid", Approve)
		// 拒绝加入
		api.Post("/refuse/:gid", Refuse)
		// 提升为管理员
		api.Post("/promote/:gid/:mid", Promote)
		// 转让群主
		api.Post("/transfer/:gid/:mid", Transfer)
		// 踢出群组
		api.Post("/fire/:gid/:mid", Fire)
		// 退出群组
		api.Post("/quit/:gid", Quit)

		// 修改群组资料
		api.Put("/edit/name/:gid", UpdateName)
		api.Put("/edit/notice/:gid", UpdateNotice)
		api.Put("/edit/addr/:gid", UpdateAddr)

		// 删除群组
		api.Delete("/remove/:gid", Remove)
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
