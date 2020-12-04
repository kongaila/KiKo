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
	user := domain.TbUser{}
	if err := l.Ctx.ReadJSON(&user); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	token, ok := l.Service.UserDetail(&user)
	if !ok {
		result = vo.Req204RespVO(0, "账号或者密码错误", nil)
		return
	}
	result = vo.Req200RespVO(0, "登录成功", token)
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
		result = vo.Req204RespVO(0, err.Error(), nil)
		return
	}
	result = vo.Req200RespVO(1, "账号创建成功", uuid)
	return
}
