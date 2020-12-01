package main

import (
	"QiqiLike/conf"
	"QiqiLike/routers"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	app := iris.New()
	// 配置路由
	routers.Router(app)
	if err := app.Run(iris.Addr(conf.Viper.GetString("server.address"))); err != nil {
		log.Println("服务退出！")
	}
}
