package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"strings"
)

type ArticleController struct {
	AttrArticleService  service.ArticleService
	AttrLoginService    service.LoginService
	AttrClubService     service.ClubService
	AttrUserClubService service.UserClubService
	AttrUserService     service.UserService
	Ctx                 iris.Context
}

// 新建一条帖子
func (a *ArticleController) PostCreate() (result *vo.RespVO) {
	article := domain.TbArticle{}
	if err := a.Ctx.ReadJSON(&article); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	if article.UserUuid, _, _ = tools.ParseHeaderToken(a.Ctx); article.UserUuid == "" {
		result = vo.Req204RespVO(1, "登录信息有误", nil)
		return
	}
	if ok := a.AttrArticleService.Create(&article); !ok {
		result = vo.Req204RespVO(1, "添加失败", nil)
		return
	}
	a.AttrUserService.UserUpdateSer(article.UserUuid, "post_num", gorm.Expr("post_num + ?", 1))
	result = vo.Req200RespVO(1, "添加成功", article.Uuid)
	return
}

// 获得帖子列表
func (a *ArticleController) GetMany() (result *vo.RespVO) {
	return
}

// 获得一个帖子详细信息
func (a *ArticleController) GetBy(uuid string) (result *vo.RespVO) {
	if strings.EqualFold(uuid, "") {
		result = vo.Req204RespVO(1, "数据有误", nil)
		return
	}
	article, err := a.AttrArticleService.GetDetailSer(uuid)
	if err != nil {
		result = vo.Req500RespVO(1, "服务器错误", nil)
		return
	}
	result = vo.Req200RespVO(1, "查询成功", article)
	return
}
