package controllers

import (
	"KiKo/datamodels/vo"
	"KiKo/service"
	"github.com/kataras/iris/v12"
)

type DictController struct {
	AttrDictService service.DictService
	Ctx             iris.Context
}

// 通过code获得字典值内容
func (d *DictController) GetCodeBy(code string) (result *vo.RespVO) {
	dicts := d.AttrDictService.GetByCodeSer(code)
	result = vo.Req200RespVO(len(dicts), "查询成功", dicts)
	return
}
