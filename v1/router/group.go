package router

import (
	"axe/gl"
	"axe/v1/system"

	"github.com/gofiber/fiber/v2"
)

func init() {
	routerNoCheck = append(routerNoCheck, registerGroupRouter)
}

func registerGroupRouter(app *fiber.App) {
	groupsAPI := fiber.New()
	gl.App.Mount("/groups", groupsAPI)
	{
		groupsAPI.Get("/{gid:int}", system.Group.GetGroupById)
		groupsAPI.Get("/", system.Group.GetGroups)
		groupsAPI.Get("/{gid:int}/pending", system.Group.GetApplyList)
		groupsAPI.Get("/{gid:int}/activities", system.Group.GetActivitiesByGroupId)

		groupsAPI.Post("/", system.Group.Create)
		groupsAPI.Post("/{gid:int}/gl.Apply", system.Group.Apply)
		groupsAPI.Post("/{gid:int}/gl.Approve", system.Group.Approve)
		groupsAPI.Post("/{gid:int}/promote/:mid", system.Group.Promote)
		groupsAPI.Post("/{gid:int}/transfer/:mid", system.Group.Transfer)
		groupsAPI.Post("/{gid:int}/remove/:mid", system.Group.Remove)
		groupsAPI.Post("/{gid:int}/quit", system.Group.Quit)

		groupsAPI.Put("/{gid:int}", system.Group.Update)
	}
}
