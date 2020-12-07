package routers

import (
	"QiqiLike/conf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Router(app *iris.Application) {
	app.Get("/ping", pong)

	// 整体路径配置
	api := app.Party(conf.Viper.GetString("api.prefix"))
	{
		// 登录注册模块路由配置
		mvc.Configure(api.Party("/"), login)

		// 用户模块
		//routerUser(api)

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
