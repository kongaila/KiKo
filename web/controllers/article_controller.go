package controllers

import (
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"github.com/kataras/iris/v12"
)

type ArticleController struct {
	AttrLoginService    service.LoginService
	AttrClubService     service.ClubService
	AttrUserClubService service.UserClubService
	Ctx                 iris.Context
}

// 创建贴吧
func (a *ArticleController) PostCreate() (result *vo.RespVO) {
	return
}

// 获得贴吧列表
func (a *ArticleController) GetMany() (result *vo.RespVO) {
	return
}

// 获得一个贴吧详情
func (a *ArticleController) GetDetail() (result *vo.RespVO) {
	return
}

func PostCreate() (result *vo.RespVO) {
	return
}
