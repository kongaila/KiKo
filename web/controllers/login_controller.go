package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"QiqiLike/tools/redis"
	"fmt"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	AttrLoginService service.LoginService
	AttrClubService  service.ClubService
	Ctx              iris.Context
}

// 登录
func (l *LoginController) PostLogin() (result *vo.RespVO) {
	user := domain.TbUser{}
	if err := l.Ctx.ReadJSON(&user); err != nil {
		fmt.Println(err)
		//panic(err)
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	token, ok := l.AttrLoginService.UserDetail(&user)
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
	uuid, err := l.AttrLoginService.Create(&user)
	if err != nil {
		result = vo.Req204RespVO(0, err.Error(), nil)
		return
	}
	result = vo.Req200RespVO(1, "账号创建成功", uuid)
	return
}

// 退出，在redis中将token删除了
func (l *LoginController) PostExit() (result *vo.RespVO) {
	token := tools.GetHeaderToken(l.Ctx)
	_ = redis.Del(token)
	result = vo.Req200RespVO(1, "退出成功", "")
	return
}
