package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerCheck = append(routerCheck, registerActivityRouter)
}

func registerActivityRouter(r fiber.Router) {
	api := r.Group("/activities")
	{
		// 获取活动详细内容
		api.Get("/:aid", handler.Activity.GetActivityByID)
		// 获取最近一段的活动
		api.Get("/list/:page/:num", handler.Activity.GetActivities)
		// TODO: 获取最近一段时间附近的活动

		// 创建活动
		api.Post("/create", handler.Activity.Create)
		// 手动终止活动
		api.Post("/end/:aid", handler.Activity.End)
		// 申请参加活动
		api.Post("/apply/:aid", handler.Activity.Apply)
		// 取消参加活动
		api.Post("/cancel/:aid", handler.Activity.Cancel)

		// 修改活动内容
		api.Put("/:aid}", handler.Activity.Update)

		// 删除活动
		api.Delete("/remove/:aid", handler.Activity.Remove)
	}
}
