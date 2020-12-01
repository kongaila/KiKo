package routers

import (
	"github.com/kataras/iris/v12"
)

func Router(app *iris.Application) {
	app.Get("/ping", pong)

	// 整体路径配置
	api := app.Party("/api/v1")
	{
		// 用户模块路由配置
		routerUser(api)
		// 贴吧模块
		// 文章模块
		// 个人中心模块
		// 热榜模块
	}
}

// 测试ping
func pong(ctx iris.Context) {
	_, _ = ctx.Write([]byte("pong!"))
}
