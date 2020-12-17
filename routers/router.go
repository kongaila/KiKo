package routers

import (
	cs "QiqiLike/constants"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Router(app *iris.Application) {
	app.Get("/ping", pong)

	// 整体路径配置
	api := app.Party(cs.Prefix)
	{
		// 前台
		mvc.Configure(api.Party("/"), login)
		mvc.Configure(api.Party("/club"), club)
		mvc.Configure(api.Party("/article"), article)
		mvc.Configure(api.Party("/user"), user)
		mvc.Configure(api.Party("/like"), like)
		mvc.Configure(api.Party("/comment"), comment)
		mvc.Configure(api.Party("/search"), search)
		mvc.Configure(api.Party("/top"), top)
		mvc.Configure(api.Party("/bulletin"), bulletin)

		// 后台
		api = api.Party("/admin")
		{
			mvc.Configure(api.Party("/admin"), admin)

		}
	}

}

// 测试ping
func pong(ctx iris.Context) {
	_, _ = ctx.Write([]byte("pong!"))
}
