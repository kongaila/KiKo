package controllers

import (
	"KiKo/service"
	"github.com/kataras/iris/v12"
)

type LikeController struct {
	AttrLikeService service.LikeService
	Ctx             iris.Context
}
