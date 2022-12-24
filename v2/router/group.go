package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerGroupRouter)
}

func registerGroupRouter(app *fiber.App) {
	// v2 版本路由
	groupsAPI := app.Group("/v2/groups")
	{
		// 获取群组详细信息
		groupsAPI.Get("/:gid", handler.Group.GetGroupByID)
		// 获取最近创建的群组
		groupsAPI.Get("/list/:page/:num", handler.Group.GetGroups)
		// 获取申请加入群组的列表
		groupsAPI.Get("/applications/:gid", handler.Group.GetApplyList)
		// 获取群组活动列表
		groupsAPI.Get("/activities/:gid", handler.Group.GetActivitiesByGroupId)

		// 创建群组
		groupsAPI.Post("/", handler.Group.Create)
		// 申请加入群组
		groupsAPI.Post("/apply/:gid/:uid", handler.Group.Apply)
		// 同意加入
		groupsAPI.Post("/approve/:gid", handler.Group.Approve)
		// 拒绝加入
		groupsAPI.Post("/refuse/:gid", handler.Group.Refuse)
		// 提升为管理员
		groupsAPI.Post("/promote/:gid/:mid", handler.Group.Promote)
		// 转让群主
		groupsAPI.Post("/transfer/:gid/:mid", handler.Group.Transfer)
		// 删除群组
		groupsAPI.Post("/remove/:gid/:mid", handler.Group.Remove)
		// 退出群组
		groupsAPI.Post("/quit/:gid", handler.Group.Quit)

		// 修改群组资料
		groupsAPI.Put("/:gid", handler.Group.Update)
	}
}
