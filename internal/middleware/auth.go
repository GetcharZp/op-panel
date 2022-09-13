package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"op-panel/helper"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		err := helper.ParseToken(token)
		if err != nil {
			return c.JSON(http.StatusOK, echo.Map{
				"code": 50014,
				"msg":  "身份认证不通过",
			})
		}
		return next(c)
	}
}
