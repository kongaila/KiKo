package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Service service.LoginService
	Ctx     iris.Context
}

// 登录
func (l *LoginController) PostLogin() (result *vo.RespVO) {
	result = vo.Req200RespVO(0, "登录", map[int]domain.TbUser{
		1: {
			Id:   1,
			Uuid: "fafasffasf",
		},
		2: {
			Id:   2,
			Uuid: "fasfasf",
		}},
	)
	return
}

// 注册
func (l *LoginController) PostRegister() (result *vo.RespVO) {
	user := domain.TbUser{}

	if err := l.Ctx.ReadJSON(&user); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
	}

	if ok := user.CheckUserNameAndPass(); !ok {
		result = vo.Req204RespVO(0, "账号不符合规则", nil)
		return
	}
	uuid, err := l.Service.Create(&user)
	if err != nil {
		result = vo.Req204RespVO(0, "账号创建失败", nil)
	}
	result = vo.Req200RespVO(1, "账号创建成功", uuid)
	return
}
