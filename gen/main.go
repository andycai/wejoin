package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./v2/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/activities?charset=utf8mb4&parseTime=true&loc=Local"))
	g.UseDB(gormdb)

	g.ApplyBasic(g.GenerateAllTable()...)

	//g.ApplyInterface(func(querier Querier) {}, model.User{}, model.Group{}, model.Activity{})

	g.Execute()
}
