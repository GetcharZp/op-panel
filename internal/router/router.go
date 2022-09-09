package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"op-panel/middleware"
	service2 "op-panel/service"
)

func Router(v1 *echo.Group) {
	v1.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	v1.POST("/login", service2.Login)

	// 需要认证操作的分组
	v2 := v1.Group("/sys", middleware.Auth)
	// 修改系统配置
	v2.PUT("/systemConfig", service2.UpdateSystemConfig)
	// 系统状态
	v2.GET("/systemState", service2.SystemState)

	// 任务管理
	v2.GET("/taskList", service2.TaskList)
	v2.POST("/taskAdd", service2.TaskAdd)
	v2.PUT("/taskEdit", service2.TaskEdit)
	v2.DELETE("/taskDelete", service2.TaskDelete)

	// 软件管理
	v2.GET("/softList", service2.SoftList)
	v2.POST("/softOperation", service2.SoftOperation)
}
