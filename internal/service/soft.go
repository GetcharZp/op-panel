package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func SoftOperation(c echo.Context) error {
	var (
		op        = c.FormValue("op")
		id        = c.FormValue("id")
		sb        = new(models.SoftBasic)
		shellPath string
		logPath   = define.LogDir + "/" + helper.GetUUID() + ".log"
	)
	if op == "" || id == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "必填参未填",
		})
	}
	err := models.DB.Model(new(models.SoftBasic)).Where("id = ?", id).First(sb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "软件未找到",
			})
		}
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	switch op {
	case "install":
		shellPath = sb.ShellInstall
	case "uninstall":
		shellPath = sb.ShellUninstall
	case "start":
		shellPath = sb.ShellStart
	case "stop":
		shellPath = sb.ShellStop
	case "restart":
		shellPath = sb.ShellRestart
	default:
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "不支持的操作",
		})
	}
	go func() {
		eq := &models.ExecuteQueue{
			SoftId:    int(sb.ID),
			State:     1,
			LogPath:   logPath,
			ShellPath: shellPath,
		}
		err = models.DB.Create(eq).Error
		if err != nil {
			log.Fatalln("[DB ERROR] : " + err.Error())
		}
		helper.RunShell(shellPath, logPath)
		err = models.DB.Model(new(models.ExecuteQueue)).Where("id = ?", eq.ID).Update("state", 2).Error
		if err != nil {
			log.Println("[DB ERROR] : " + err.Error())
		}
	}()
	return c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "请求成功",
	})
}
