package controllers

import (
	"KiKo/datamodels/domain"
	"KiKo/datamodels/vo"
	"KiKo/service"
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

// 创建公告
func (b *BulletinController) Post() (result *vo.RespVO) {
	tb := domain.TbBulletin{}
	if err := b.Ctx.ReadJSON(&tb); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	if ok := b.AttrBulletinService.Create(tb); !ok {
		result = vo.Req204RespVO(1, "添加失败", nil)
		return
	}
	result = vo.Req200RespVO(1, "添加成功", nil)
	return
}

// 获得历史公告
func (b *BulletinController) GetMany() (result *vo.AdminRespVO) {
	count, data := b.AttrBulletinService.GetManySer(b.Ctx.URLParams())
	result = vo.ReqSuccessAdminRespVO(count, "查询成功", data)
	return
}
