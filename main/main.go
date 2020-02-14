package main

import (
	"github.com/kataras/iris"
)
func main(){
	app := iris.New()

	app.Get("/getRequest", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.HTML("请求的路径是:" + path)
	})
	app.Post("/login",func(context iris.Context){
		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info("用户名是： " + name + "密码是: " + pwd)

	})
	var person Person
	app.Post("/getjson",func(context iris.Context){
		if err := context.ReadJSON(&person); err != nil{
			panic(err)
		}
		context.Writef("Received: %#+v\n", person)

	})
	app.Handle("GET","/weather/{date}/{city}",func(context iris.Context){
		date :=  context.Params().Get("date")
		city :=  context.Params().Get("city")
		context.Writef("日期是： :" + date + "城市是：" + city )
	})


	app.Run(iris.Addr(":8080"))
}

type Person struct {
	Name string
	Age int
}tia

type Weather struct {
	Date string
	City string
}