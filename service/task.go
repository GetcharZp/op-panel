package service

import (
	"bufio"
	"errors"
	"github.com/kataras/iris/v12"
	"log"
	"op-panel/define"
	"op-panel/helper"
	"op-panel/models"
	"os"
	"path"
	"syscall"
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

func TaskAdd(c iris.Context) {
	name := c.PostValue("name")
	spec := c.PostValue("spec")
	data := c.PostValue("data")
	if name == "" || spec == "" || data == "" {
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
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
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
		return
	}
	f, err := os.Create(tb.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(tb.ShellPath), 0777)
		f, err = os.Create(tb.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			c.JSON(iris.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
			return
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(iris.Map{
		"code": 200,
		"msg":  "新增成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
}

func TaskEdit(c iris.Context) {
	id := c.PostValue("id")
	name := c.PostValue("name")
	spec := c.PostValue("spec")
	data := c.PostValue("data")
	if id == "" || name == "" || spec == "" || data == "" {
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}

	tb := new(models.TaskBasic)
	err := models.DB.Where("id = ?", id).First(tb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
		return
	}
	tb.Name = name
	tb.Spec = spec
	err = models.DB.Where("id = ?", id).Updates(tb).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
		return
	}

	f, err := os.Create(tb.ShellPath)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(path.Dir(tb.ShellPath), 0777)
		f, err = os.Create(tb.ShellPath)
		if err != nil {
			log.Println("[CREATE FILE ERROR] : " + err.Error())
			c.JSON(iris.Map{
				"code": -1,
				"msg":  "系统异常 : " + err.Error(),
			})
			return
		}
	}
	w := bufio.NewWriter(f)
	w.WriteString(data)
	w.Flush()

	c.JSON(iris.Map{
		"code": 200,
		"msg":  "编辑成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
}

func TaskDelete(c iris.Context) {
	id := c.PostValue("id")
	err := models.DB.Where("id = ?", id).Delete(new(models.TaskBasic)).Error
	if err != nil {
		log.Println("[DB ERROR] : " + err.Error())
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常 : " + err.Error(),
		})
		return
	}

	c.JSON(iris.Map{
		"code": 200,
		"msg":  "删除成功",
	})

	syscall.Kill(define.PID, syscall.SIGINT)
}
