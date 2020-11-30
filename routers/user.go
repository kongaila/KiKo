package routers

import "github.com/kataras/iris/v12/core/router"

func routerUser(party router.Party) {
	party.Post("/login", nil)
}
