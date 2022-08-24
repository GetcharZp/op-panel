package main

import (
	stdContext "context"
	"fmt"
	"github.com/kataras/iris/v12"
	"op-panel/define"
	"op-panel/models"
	"op-panel/service"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	define.PID = syscall.Getpid()
	fmt.Println("PID : " + strconv.Itoa(define.PID))
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
	v2.Put("/systemConfig", service.UpdateSystemConfig)

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
