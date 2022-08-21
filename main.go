package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"op-panel/models"
	"op-panel/service"
)

func main() {
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
	app.Listen(sc.Port)
}
