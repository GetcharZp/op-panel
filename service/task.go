package service

import (
	"github.com/kataras/iris/v12"
	"log"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
)

func TaskList(c iris.Context) {
	var (
		index, _ = c.URLParamInt("index")
		size, _  = c.URLParamInt("size")
		tb       = make([]*models.TaskBasic, 0)
		cnt      int64
	)

	index = helper.If(index == 0, 1, index).(int)
	size = helper.If(size == 0, define.PageSize, size).(int)

	err := models.DB.Model(new(models.TaskBasic)).Count(&cnt).Offset(index).Limit(size).Find(&tb).Error
	if err != nil {
		log.Println("[DB ERROR]" + err.Error())
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(iris.Map{
		"code": 200,
		"msg":  "加载成功",
		"data": iris.Map{
			"list":  tb,
			"count": cnt,
		},
	})
}
