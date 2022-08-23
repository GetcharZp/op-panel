package main

import (
	stdContext "context"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"op-panel/define"
	"op-panel/models"
	"op-panel/service"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var PID int

func main() {
	PID := syscall.Getpid()
	fmt.Println("PID : " + strconv.Itoa(PID))
	models.NewDB()
	sc := service.GetSystemConfig()
	ub := service.InitUserConfig()
	fmt.Println("Address : http://localhost" + sc.Port + sc.Entry)
	fmt.Println("Username : " + ub.Name)
	fmt.Println("Password : " + ub.Password)
	app := iris.New()

	v1 := app.Party(sc.Entry)
	v1.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	v1.Post("/login", service.Login)

	// 需要认证操作的分组
	v2 := v1.Party("/sys")
	// 修改系统配置
	v2.Put("/systemConfig", UpdateSystemConfig)

	run := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)
		select {
		case <-ch:
			timeout := 10 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
			go main()
		}
	}()
	app.Listen(sc.Port, iris.WithoutInterruptHandler)
	<-run
}

// UpdateSystemConfig 修改系统配置
func UpdateSystemConfig(c iris.Context) {
	var (
		port  = c.PostValue("port")
		entry = c.PostValue("entry")
		cb    = new(models.ConfigBasic)
		sc    = new(define.SystemConfig)
	)
	// 获取现用配置
	err := models.DB.Where("key = 'system'").First(cb).Error
	if err != nil {
		log.Printf("[DB ERROR] : %v\n", err)
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}
	json.Unmarshal([]byte(cb.Value), sc)

	// 设置新的配置
	if port != "" {
		sc.Port = ":" + port
	}
	if entry != "" {
		sc.Entry = "/" + entry
	}
	scByte, _ := json.Marshal(sc)

	// 更新配置
	err = models.DB.Model(new(models.ConfigBasic)).Where("key = 'system'").Update("value", string(scByte)).Error
	if err != nil {
		log.Printf("[DB ERROR] : %v\n", err)
		c.JSON(iris.Map{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}

	c.JSON(iris.Map{
		"code": 200,
		"msg":  "修改成功",
	})

	// 重启服务
	syscall.Kill(PID, syscall.SIGINT)
}
