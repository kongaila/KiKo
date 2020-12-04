package main

import (
	"QiqiLike/conf"
	"QiqiLike/routers"
	"github.com/kataras/iris/v12"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := iris.New()
	// 配置路由
	routers.Router(app)
	go func() {
		log.Printf(" [INFO] 项目启动到端口%s\n", conf.Viper.GetString("server.address"))
		if err := app.Run(iris.Addr(conf.Viper.GetString("server.address"))); err != nil {
			log.Fatalf(" [ERROR] 项目启动到地址:%s err:%v\n", conf.Viper.GetString("server.address"), err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 优雅的退出
	<-quit
}
