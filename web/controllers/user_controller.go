package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/kataras/iris/v12"
	"strings"
)

type UserController struct {
	AttrUserService service.UserService
	AttrLikeService service.LikeService
	Ctx             iris.Context
}

// 加入收藏 TODO 添加uuid不存在问题待解决
func (u *UserController) PostLikeBy(uuid string) (result *vo.RespVO) {
	if strings.EqualFold(uuid, "") {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
	like := domain.TbLike{
		UserUuid:    userUuid,
		ArticleUuid: uuid,
		Title:       u.Ctx.PostValue("title"),
	}
	uuid, err := u.AttrLikeService.CreateLike(like)
	if err != nil {
		result = vo.Req204RespVO(0, err.Error(), nil)
		return
	}
	result = vo.Req200RespVO(1, "加入成功", uuid)
	return
}
