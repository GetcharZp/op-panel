package service

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
)

func Login(c iris.Context) {
	name := c.PostValue("name")
	password := c.PostValue("password")
	if name == "" || password == "" {
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	cb := new(models.ConfigBasic)
	err := models.DB.Model(new(models.ConfigBasic)).Where("key = 'user'").
		First(cb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(iris.Map{
				"code": -1,
				"msg":  "用户信息未初始化",
			})
			return
		}
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}
	ub := new(define.UserBasic)
	json.Unmarshal([]byte(cb.Value), ub)
	if ub.Password != password || ub.Name != name {
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "用户名或密码不正确",
		})
		return
	}
	token, err := helper.GenerateToken()
	if err != nil {
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}
	c.JSON(iris.Map{
		"code": 200,
		"data": iris.Map{
			"token": token,
		},
		"msg": "登录成功",
	})
}
