package service

import (
	"bufio"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
	"os"
	"path"
	"strconv"
	"syscall"
)

func TaskList(c echo.Context) error {
	var (
		index, _ = strconv.Atoi(c.QueryParam("index"))
		size, _  = strconv.Atoi(c.QueryParam("size"))
		tb       = make([]*models.TaskBasic, 0)
		cnt      int64
	)

	size = helper.If(size == 0, define.PageSize, size).(int)
	index = helper.If(index == 0, 1, index).(int)

	err := models.DB.Model(new(models.TaskBasic)).Count(&cnt).Offset((1 - index) * size).Limit(size).Find(&tb).Error
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
			"list":  tb,
			"count": cnt,
		},
	})
}

func TaskAdd(c echo.Context) error {
	name := c.FormValue("name")
	spec := c.FormValue("spec")
	data := c.FormValue("data")
	if name == "" || spec == "" || data == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "必填参不能为空",
		})
	}
	shellName := helper.GetUUID()
	tb := &models.TaskBasic{
		Name:      name,
		Spec:      spec,
		ShellPath: define.ShellDir + "/" + shellName + ".sh",
		LogPath:   define.LogDir + "/" + shellName + ".log",
	}
	err := models.DB.Create(tb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	f, err := os.Create(tb.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(tb.ShellPath), 0777)
		f, err = os.Create(tb.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "新增成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

func TaskEdit(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	spec := c.FormValue("spec")
	data := c.FormValue("data")
	if id == "" || name == "" || spec == "" || data == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "必填参不能为空",
		})
	}

	tb := new(models.TaskBasic)
	err := models.DB.Where("id = ?", id).First(tb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}
	tb.Name = name
	tb.Spec = spec
	err = models.DB.Where("id = ?", id).Updates(tb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}

	f, err := os.Create(tb.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(tb.ShellPath), 0777)
		f, err = os.Create(tb.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			return c.JSON(http.StatusOK, echo.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "编辑成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
	return nil
}

func TaskDelete(c echo.Context) error {
	id := c.FormValue("id")
	err := models.DB.Where("id = ?", id).Delete(new(models.TaskBasic)).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		return c.JSON(http.StatusOK, echo.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, echo.Map{
		"code": 200,
		"msg":  "删除成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)

	return nil
}
