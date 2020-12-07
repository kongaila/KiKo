package middleware

import (
	cs "QiqiLike/constants"
	"github.com/kataras/iris/v12"
	"log"
)

// 打印日志
func message(ctx iris.Context) {
	log.Printf(cs.INFO+"发送了请求： %s %s%s", ctx.Method(), ctx.Host(), ctx.Request().URL.String())
	ctx.Next()
}

// 校验token
func checkToken(ctx iris.Context) {
	log.Println(cs.INFO, "校验token!")
	// TODO
	ctx.Next()
}

// 最后执行（目前没有用到）
func after(ctx iris.Context) {
	log.Println(cs.INFO, "走完了!")
}
