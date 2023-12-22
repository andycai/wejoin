package group

import "github.com/andycai/axe-fiber/model"

type Result struct {
	Group       model.Group
	Member      model.GroupMember
	Application model.GroupApplication
	Error       error
}
