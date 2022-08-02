package controllers

import (
	"KiKo/datamodels/domain"
	"KiKo/datamodels/vo"
	"KiKo/service"
	"github.com/kataras/iris/v12"
	"log"
)

type AdminController struct {
	AttrArticleService  service.ArticleService
	AttrLoginService    service.LoginService
	AttrClubService     service.ClubService
	AttrUserClubService service.UserClubService
	AttrUserService     service.UserService
	AttrBulletinService service.BulletinService
	AttrCommentService  service.CommentService
	AttrAdminService    service.AdminService
	Ctx                 iris.Context
}

// 登录
func (a *AdminController) PostLogin() (result *vo.AdminRespVO) {
	admin := domain.TbAdmin{}
	if err := a.Ctx.ReadBody(&admin); err != nil {
		log.Println(err)
		//panic(err)
		result = vo.ReqFailedAdminRespVO(0, "数据有误", nil)
		return
	}
	token, ok := a.AttrAdminService.AdminDetail(&admin)
	if !ok {
		result = vo.ReqFailedAdminRespVO(0, "账号或者密码错误", nil)
		return
	}
	result = vo.ReqSuccessAdminRespVO(0, "登录成功", token)
	return
}

// 查询帖吧
func (a *AdminController) GetClub() (result *vo.AdminRespVO) {
	data, count, err := a.AttrClubService.GetClubManySer(a.Ctx.URLParams())
	if err != nil {
		result = vo.ReqFailedAdminRespVO(count, err.Error(), nil)
		return
	}
	result = vo.ReqSuccessAdminRespVO(count, "查询成功", data)
	return
}

// 查询帖子
func (a *AdminController) GetArticle() (result *vo.AdminRespVO) {
	data, count, err := a.AttrArticleService.GetArticleManySer(a.Ctx.URLParams())
	if err != nil {
		result = vo.ReqFailedAdminRespVO(count, err.Error(), nil)
		return
	}
	result = vo.ReqSuccessAdminRespVO(count, "查询成功", data)
	return
}

// 查询用户
func (a *AdminController) GetUser() (result *vo.AdminRespVO) {
	data, count, err := a.AttrUserService.GetUserManySer(a.Ctx.URLParams())
	if err != nil {
		result = vo.ReqFailedAdminRespVO(count, err.Error(), nil)
		return
	}
	result = vo.ReqSuccessAdminRespVO(count, "查询成功", data)
	return
}

// 查询评论
func (a *AdminController) GetComment() (result *vo.AdminRespVO) {
	comments, count := a.AttrCommentService.GetCommentManySer(a.Ctx.URLParams())
	result = vo.ReqSuccessAdminRespVO(count, "查询成功", comments)
	return
}

// 审核
func (a *AdminController) PostCheck() (result *vo.AdminRespVO) {
	uuid := a.Ctx.PostValue("uuid")
	i, _ := a.Ctx.PostValueInt("type")
	var msg string
	if i == 1 {
		// 通过
		a.AttrClubService.UpdateClubInfoSer(uuid, "status", "1")
		msg = "通过成功！"
	} else {
		// 不通过
		a.AttrClubService.UpdateClubInfoSer(uuid, "status", "2")
		msg = "驳回成功！"
	}
	result = vo.ReqSuccessAdminRespVO(0, msg, nil)
	return
}

// 举报审核
func (a *AdminController) PostReport() (result *vo.AdminRespVO) {
	uuid := a.Ctx.PostValue("uuid")
	i, _ := a.Ctx.PostValueInt("type")
	var msg string
	if i == 1 {
		// 通过
		a.AttrArticleService.UpdateArticleInfoSer(uuid, "status", "2")
		msg = "通过成功！"
	} else {
		// 不通过
		a.AttrArticleService.UpdateArticleInfoSer(uuid, "status", "3")
		msg = "驳回成功！"
	}
	result = vo.ReqSuccessAdminRespVO(0, msg, nil)
	return
}
