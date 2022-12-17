package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerActivityRouter)
}

func registerActivityRouter(app *fiber.App) {
	// v2 版本路由
	actsAPI := app.Group("/v2/activities")
	{
		// 获取活动详细内容
		actsAPI.Get("/:aid", handler.Activity.GetActivityById)
		// 获取最近一段的活动
		actsAPI.Get("/", handler.Activity.GetActivities)
		// TODO: 获取最近一段时间附近的活动

		// 创建活动
		actsAPI.Post("/", handler.Activity.Create)
		// 手动终止活动
		actsAPI.Post("/:aid/end", handler.Activity.End)
		// 申请参加活动
		actsAPI.Post("/:aid/apply", handler.Activity.Apply)
		// 取消参加活动
		actsAPI.Post("/:aid/cancel", handler.Activity.Cancel)
		// 删除活动
		actsAPI.Post("/:aid/remove/:index", handler.Activity.Remove)

		// 修改活动内容
		actsAPI.Put("/:aid}", handler.Activity.Update)
	}
}
