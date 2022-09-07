package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
	"strconv"
)

func SoftList(c echo.Context) error {
	var (
		index, _ = strconv.Atoi(c.QueryParam("index"))
		size, _  = strconv.Atoi(c.QueryParam("size"))
		sb       = make([]*models.SoftBasic, 0)
		cnt      int64
	)

	size = helper.If(size == 0, define.PageSize, size).(int)
	index = helper.If(index == 0, 1, index).(int)

	err := models.DB.Model(new(models.SoftBasic)).Count(&cnt).Offset((1 - index) * size).Limit(size).Find(&sb).Error
	if err != nil {
		log.Println("[DB ERROR]" + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "加载成功",
		"data": echo.Map{
			"list":  sb,
			"count": cnt,
		},
	})
}
