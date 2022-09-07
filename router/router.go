package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"op-panel/middleware"
	"op-panel/service"
)

func Router(v1 *echo.Group) {
	v1.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	v1.POST("/login", service.Login)

	// 需要认证操作的分组
	v2 := v1.Group("/sys", middleware.Auth)
	// 修改系统配置
	v2.PUT("/systemConfig", service.UpdateSystemConfig)
	// 系统状态
	v2.GET("/systemState", service.SystemState)

	// 任务管理
	v2.GET("/taskList", service.TaskList)
	v2.POST("/taskAdd", service.TaskAdd)
	v2.PUT("/taskEdit", service.TaskEdit)
	v2.DELETE("/taskDelete", service.TaskDelete)

	// 软件管理
	v2.GET("/softList", service.SoftList)
}
