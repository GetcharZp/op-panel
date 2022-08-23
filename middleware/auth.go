package middleware

import (
	"github.com/kataras/iris/v12"
	"op-panel/helper"
)

func Auth() iris.Handler {
	return func(c iris.Context) {
		token := c.GetHeader("toke")
		err := helper.ParseToken(token)
		if err != nil {
			c.JSON(iris.Map{
				"code": -1,
				"msg":  "身份认证不通过",
			})
			return
		}
		c.Next()
	}
}
