package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerGroupRouter)
}

func registerGroupRouter(r fiber.Router) {
	api := r.Group("/groups")
	{
		// 获取群组详细信息
		api.Get("/:gid", handler.Group.GetGroupByID)
		// 获取最近创建的群组
		api.Get("/list/:page/:num", handler.Group.GetGroups)
		// 获取申请加入群组的列表
		api.Get("/applications/:gid", handler.Group.GetApplyList)
		// 获取群组活动列表
		api.Get("/activities/:gid/:page/:num", handler.Group.GetActivitiesByGroupId)

		// 创建群组
		api.Post("/create", handler.Group.Create)
		// 申请加入群组
		api.Post("/apply/:gid/:uid", handler.Group.Apply)
		// 同意加入
		api.Post("/approve/:gid", handler.Group.Approve)
		// 拒绝加入
		api.Post("/refuse/:gid", handler.Group.Refuse)
		// 提升为管理员
		api.Post("/promote/:gid/:mid", handler.Group.Promote)
		// 转让群主
		api.Post("/transfer/:gid/:mid", handler.Group.Transfer)
		// 踢出群组
		api.Post("/fire/:gid/:mid", handler.Group.Fire)
		// 退出群组
		api.Post("/quit/:gid", handler.Group.Quit)

		// 修改群组资料
		api.Put("/edit/name/:gid", handler.Group.UpdateName)
		api.Put("/edit/notice/:gid", handler.Group.UpdateNotice)
		api.Put("/edit/addr/:gid", handler.Group.UpdateAddr)

		// 删除群组
		api.Delete("/remove/:gid", handler.Group.Remove)
	}
}
