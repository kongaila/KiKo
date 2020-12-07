package main

import (
	"QiqiLike/conf"
	cs "QiqiLike/constants"
	"QiqiLike/routers"
	"QiqiLike/web/middleware"
	"github.com/kataras/iris/v12"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := iris.New()
	// 配置中间件
	middleware.RegisterMiddleware(app)
	// 配置路由
	routers.Router(app)
	go func() {
		log.Printf(cs.INFO+"项目启动到端口%s\n", conf.Viper.GetString("server.address"))
		if err := app.Run(iris.Addr(conf.Viper.GetString("server.address"))); err != nil {
			log.Fatalf(cs.ERROR+"项目启动到地址:%s err:%v\n", conf.Viper.GetString("server.address"), err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 监听退出信号
	<-quit
}
