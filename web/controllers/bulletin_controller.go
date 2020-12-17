package controllers

import (
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"github.com/kataras/iris/v12"
)

type BulletinController struct {
	AttrBulletinService service.BulletinService
	Ctx                 iris.Context
}

// 获得最新公告板内容
func (b *BulletinController) Get() (result *vo.RespVO) {
	data := b.AttrBulletinService.GetBulletinSer()
	result = vo.Req200RespVO(1, "查询成功", data)
	return
}
