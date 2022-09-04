package router

import (
	"github.com/kataras/iris/v12"
	"op-panel/service"
)

func Router(v1 iris.Party) {
	v1.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	v1.Post("/login", service.Login)

	// 需要认证操作的分组
	v2 := v1.Party("/sys")
	// 修改系统配置
	v2.Put("/systemConfig", service.UpdateSystemConfig)
	// 系统状态
	v2.Get("/systemState", service.SystemState)

	// 任务列表
	v2.Get("/taskList", service.TaskList)
	v2.Post("/taskAdd", service.TaskAdd)
}
