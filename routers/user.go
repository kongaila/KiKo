package routers

import (
	"QiqiLike/web/controllers"
	"github.com/kataras/iris/v12/core/router"
)

func routerUser(party router.Party) {
	party.Post("/register", controllers.PostRegister)
}
