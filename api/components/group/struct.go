package group

import "github.com/andycai/wejoin/model"

type Result struct {
	Group       model.Group
	Member      model.GroupMember
	Application model.GroupApplication
	Error       error
}
