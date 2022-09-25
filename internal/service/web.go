package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
	"os"
	"path"
	"strconv"
)

func WebList(c echo.Context) error {
	var (
		index, _ = strconv.Atoi(c.QueryParam("index"))
		size, _  = strconv.Atoi(c.QueryParam("size"))
		wb       = make([]*models.WebBasic, 0)
		cnt      int64
	)
	size = helper.If(size == 0, define.PageSize, size).(int)
	index = helper.If(index == 0, 1, index).(int)
	err := models.DB.Model(new(models.WebBasic)).Count(&cnt).Offset((index - 1) * size).Limit(size).Find(&wb).Error
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
		"data": wb,
	})
}

func WebAdd(c echo.Context) error {
	var (
		name   = c.FormValue("name")
		domain = c.FormValue("domain")
		cnt    int64
	)

	if name == "" || domain == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "必填参不能为空",
		})
	}
	// 判断域名是否已存在
	err := models.DB.Model(new(models.WebBasic)).Where("domain = ?", domain).Count(&cnt).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	if cnt > 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "域名已存在",
		})
	}
	wb := &models.WebBasic{
		Name:   name,
		Domain: domain,
		Dir:    define.DefaultWebDir + domain,
	}
	// 创建网站目录
	err = os.MkdirAll(path.Dir(wb.Dir), 0777)
	if err != nil {
		log.Println("[CREATE DIR ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	// 创建新的网站记录
	err = models.DB.Create(wb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}

	// TODO: 创建nginx配置文件 & 重启nginx加载配置
	return c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "新增成功",
	})
}
