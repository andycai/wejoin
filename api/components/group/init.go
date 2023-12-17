package group

import (
	"github.com/andycai/axe-fiber/core"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const (
	KeyPostDB            = "group.gorm.db"
	KeyPostNoCheckRouter = "group.router.nocheck"
	KeyPostCheckRouter   = "group.router.check"
)

var db *gorm.DB

func initDB(dbs []*gorm.DB) {
	db = dbs[0]
}

// `POST /api/questions` —— 创建一个资源
// `POST /api/questions/existed_question` —— 更新一个已存在的资源
// `POST /api/questions/new_question` —— 创建一个资源（错误的做法）

// `PUT /api/questions/existed_question`  —— 覆盖一个已存在的资源（覆盖意味着，如果缺少字段，则原有字段会被删除）
// `PUT /api/questions` —— 创建一个资源（错误的做法）
// `PUT /api/questions/new_question` —— 创建一个资源

// `PATCH /api/questions` —— 创建一个资源（错误的做法）
// `PATCH /api/questions/existed_question` —— 更新一个已存在的资源
// `PATCH /api/questions/new_question` —— 创建一个资源（错误的做法）

/*
1. 资源操作由请求方式决定（method）
- https://api.baidu.com/books      - get请求：获取所有书
- https://api.baidu.com/books/1    - get请求：获取主键为1的书
- https://api.baidu.com/books      - post请求：新增一本书书
- https://api.baidu.com/books/1    - put请求：整体修改主键为1的书
- https://api.baidu.com/books/1    - patch请求：局部修改主键为1的书
- https://api.baidu.com/books/1    - delete请求：删除主键为1的书
2. 过滤，通过在url传参的形式传递搜索条件
- https://api.example.com/v1/zoos?limit=10               指定返回记录的数量
- https://api.example.com/v1/zoos?offset=10              指定返回记录的开始位置
- https://api.example.com/v1/zoos?page=2&per_page=100    指定第几页，以及每页的记录数
- https://api.example.com/v1/zoos?sortby=name&order=asc  指定返回结果按照哪个属性排序，以及排序顺序
- https://api.example.com/v1/zoos?animal_type_id=1       指定筛选条件
3. 响应状态码
3.1 正常响应，响应状态码2xx
- 200：常规请求
- 201：创建成功
3.2 重定向响应，响应状态码3xx
- 301：永久重定向
- 302：暂时重定向
3.3 客户端异常，响应状态码4xx
- 403：请求无权限
- 404：请求路径不存在
- 405：请求方法不存在
3.4 服务器异常，响应状态码5xx
- 500：服务器异常
*/

func initNoCheckRouter(r fiber.Router) {
	api := r.Group("/v2")
	{
		// query: ?page=2
		api.Get("/groups/:gid", GetGroupByID)
		api.Get("/groups", GetGroupsByPage)
		api.Get("/groups/applications/:gid", GetApplyList)

		// json body
		api.Post("/groups", Create)
		api.Post("/groups/apply", Apply)
		api.Post("/groups/approve", Approve)
		api.Post("/groups/refuse", Refuse)
		api.Post("/groups/promote", Promote)
		api.Post("/groups/transfer", Transfer)
		api.Post("/groups/fire", Fire)
		api.Post("/groups/quit", Quit)

		// json body and url param
		api.Put("/groups/name/:gid", UpdateName)
		api.Put("/groups/notice/:gid", UpdateNotice)
		api.Put("/groups/addr/:gid", UpdateAddr)

		api.Delete("/groups/:gid", Remove)
	}
}

func initCheckRouter(r fiber.Router) {
	//
}

func init() {
	core.RegisterDatabase(KeyPostDB, initDB)
	core.RegisterNoCheckRouter(KeyPostNoCheckRouter, initNoCheckRouter)
	core.RegisterCheckRouter(KeyPostCheckRouter, initCheckRouter)
}
