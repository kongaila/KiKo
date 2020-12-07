package middleware

import (
	"QiqiLike/datamodels/vo"
	"github.com/kataras/iris/v12"
	"log"
)

// 500时候的错误， 全局异常的处理器
func error500Handler(ctx iris.Context) {
	log.Println("服务器内部错误！")
	ctx.JSON(vo.Req500RespVO(0, "服务器内部错误", nil))
}

// 400时候的错误
func error400Handler(ctx iris.Context) {
	// TODO
}
