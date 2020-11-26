package main

import (
	"QiqiLike/conf"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Run(iris.Addr(conf.Viper.GetString("server.address")))
	//get := conf.Viper.GetStringSlice("server.address")
	//fmt.Println(get)
}
