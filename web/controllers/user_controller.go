package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"github.com/kataras/iris/v12"
)

// 注册接口
func PostRegister(ctx iris.Context) {
	user := domain.TbUser{}
	_ = ctx.ReadJSON(&user)

	_, _ = ctx.JSON(vo.Req200RespVO(0, nil))
}
