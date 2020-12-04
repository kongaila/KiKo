package routers

import (
	repo "QiqiLike/repositorys"
	"QiqiLike/service"
	"QiqiLike/web/controllers"
	"github.com/kataras/iris/v12/mvc"
)

// 登录模块
func login(app *mvc.Application) {
	login := repo.NewLoginRepository(repo.Db)
	loginService := service.NewLoginService(login)
	app.Register(loginService)
	app.Handle(new(controllers.LoginController))
}
