package router

import (
	"github.com/andycai/axe-fiber/v1/system"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerGroupRouter)
}

func registerGroupRouter(app *fiber.App) {
	groupsAPI := app.Group("/groups")
	{
		groupsAPI.Get("/:gid", system.Group.GetGroupById)
		groupsAPI.Get("/", system.Group.GetGroups)
		groupsAPI.Get("/:gid/pending", system.Group.GetApplyList)
		groupsAPI.Get("/:gid/activities", system.Group.GetActivitiesByGroupId)

		groupsAPI.Post("/", system.Group.Create)
		groupsAPI.Post("/:gid/apply", system.Group.Apply)
		groupsAPI.Post("/:gid/approve", system.Group.Approve)
		groupsAPI.Post("/:gid/promote/:mid", system.Group.Promote)
		groupsAPI.Post("/:gid/transfer/:mid", system.Group.Transfer)
		groupsAPI.Post("/:gid/remove/:mid", system.Group.Remove)
		groupsAPI.Post("/:gid/quit", system.Group.Quit)

		groupsAPI.Put("/:gid", system.Group.Update)
	}
}
