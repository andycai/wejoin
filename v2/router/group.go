package router

import (
	"github.com/andycai/axe-fiber/v2/handler"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerGroupRouter)
}

func registerGroupRouter(app *fiber.App) {
	groupsAPI := app.Group("/groups")
	{
		groupsAPI.Get("/:gid", handler.Group.GetGroupById)
		groupsAPI.Get("/", handler.Group.GetGroups)
		groupsAPI.Get("/:gid/pending", handler.Group.GetApplyList)
		groupsAPI.Get("/:gid/activities", handler.Group.GetActivitiesByGroupId)

		groupsAPI.Post("/", handler.Group.Create)
		groupsAPI.Post("/:gid/apply", handler.Group.Apply)
		groupsAPI.Post("/:gid/approve", handler.Group.Approve)
		groupsAPI.Post("/:gid/promote/:mid", handler.Group.Promote)
		groupsAPI.Post("/:gid/transfer/:mid", handler.Group.Transfer)
		groupsAPI.Post("/:gid/remove/:mid", handler.Group.Remove)
		groupsAPI.Post("/:gid/quit", handler.Group.Quit)

		groupsAPI.Put("/:gid", handler.Group.Update)
	}
}
