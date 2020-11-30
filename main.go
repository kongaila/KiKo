package main

import (
	"QiqiLike/conf"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()

	if err := app.Run(iris.Addr(conf.Viper.GetString("server.address"))); err != nil {
		log.Println("服务启动失败！")
		panic(err)
	}
}
