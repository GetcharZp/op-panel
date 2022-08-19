package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"op-panel/models"
	"op-panel/service"
)

func main() {
	models.NewDB()
	sc := service.GetSystemConfig()
	log.Println("Address : http://localhost" + sc.Port + sc.Entry)
	app := iris.New()

	v1 := app.Party(sc.Entry)
	v1.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	app.Listen(sc.Port)
}
