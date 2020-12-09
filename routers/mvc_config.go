package routers

import (
	"QiqiLike/web/controllers"
	"github.com/kataras/iris/v12/mvc"
)

// 登录模块
func login(app *mvc.Application) {
	app.Register(getLoginService())
	app.Handle(new(controllers.LoginController))
}

// 贴吧模块
func club(app *mvc.Application) {
	app.Register(getLoginService(), getClubService(), getUserClubService(), getArticleService())
	app.Handle(new(controllers.ClubController))
}
