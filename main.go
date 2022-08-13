package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"code": 200,
			"msg":  "success",
		})
	})
	app.Listen(":8000")
}
