package routers

import "github.com/kataras/iris/v12"

func Router(app *iris.Application) {
	app.Get("/ping", pong)

	// 整体配置
	v1 := app.Party("/v1")
	{
		// 用户模块路由配置
		routerUser(v1)
	}
}

// 测试ping
func pong(ctx iris.Context) {
	_, _ = ctx.Write([]byte("pong!"))
}
