package service

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
)

func Login(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	if name == "" || password == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
	}
	cb := new(models.ConfigBasic)
	err := models.DB.Model(new(models.ConfigBasic)).Where("key = 'user'").
		First(cb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "用户信息未初始化",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}
	ub := new(define.UserBasic)
	json.Unmarshal([]byte(cb.Value), ub)
	if ub.Password != password || ub.Name != name {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "用户名或密码不正确",
		})
	}
	token, err := helper.GenerateToken()
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"data": echo.Map{
			"token": token,
		},
		"msg": "登录成功",
	})
}
