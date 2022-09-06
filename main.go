package main

import (
	stdContext "context"
	"fmt"
	"github.com/labstack/echo/v4"
	"op-panel/define"
	"op-panel/models"
	"op-panel/router"
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

	// 定时任务
	cron := make(chan struct{})
	go service.Cron(cron)

	e := echo.New()
	v1 := e.Group(sc.Entry)
	router.Router(v1)

	run := make(chan struct{})
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)
		select {
		case <-ch:
			timeout := 10 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			e.Shutdown(ctx)
			cron <- struct{}{}
			go main()
		}
	}()
	e.Logger.Print(e.Start(sc.Port))
	<-run
}
