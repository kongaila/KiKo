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
	AttrUserService     service.UserService
	AttrLikeService     service.LikeService
	AttrArticleService  service.ArticleService
	AttrClubService     service.ClubService
	AttrCommentService  service.CommentService
	AttrUserClubService service.UserClubService
	Ctx                 iris.Context
}

// 查看是否加入了贴吧
func (u *UserController) PostJoinBy(clubUuid string) (result *vo.RespVO) {
	userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
	ok := u.AttrUserClubService.GetIsJoinClub(userUuid, clubUuid)
	result = vo.Req200RespVO(0, "查询成功", ok)
	return
}

// 加入收藏
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

// 查看自己（别人）的详细信息
func (u *UserController) PostDetail() (result *vo.RespVO) {
	var uuid string
	if uuid, _, _ = tools.ParseHeaderToken(u.Ctx); uuid == "" {
		result = vo.Req204RespVO(0, "没有找到登录信息", nil)
		return
	}
	userUuid := u.Ctx.PostValue("userUuid")
	if !strings.EqualFold(userUuid, "") {
		uuid = userUuid
	}
	user := u.AttrUserService.GetUserDetailSer(uuid)
	result = vo.Req200RespVO(1, "查询成功", user)
	return
}

// 查看自己（别人）发布的帖子
func (u *UserController) GetArticle() (result *vo.RespVO) {
	params := u.Ctx.URLParams()
	if _, ok := params["userUuid"]; !ok {
		userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
		params["userUuid"] = userUuid
	}
	params["flag"] = "true"
	articles, count, err := u.AttrArticleService.GetArticleManySer(params)
	if err != nil {
		result = vo.Req204RespVO(0, "查询错误", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", articles)
	return
}

// 查看自己（别人）的评论
func (u *UserController) GetComment() (result *vo.RespVO) {
	params := u.Ctx.URLParams()
	if _, ok := params["userUuid"]; !ok {
		userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
		params["userUuid"] = userUuid
	}
	params["flag"] = "true"
	comments, count := u.AttrCommentService.GetCommentManySer(params)
	result = vo.Req200RespVO(count, "查询成功", comments)
	return
}

// 查看自己（别人）收藏的帖子
func (u *UserController) GetLikeArticle() (result *vo.RespVO) {
	params := u.Ctx.URLParams()
	if _, ok := params["userUuid"]; !ok {
		userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
		params["userUuid"] = userUuid
	}
	likes, count, err := u.AttrLikeService.GetLikeManySer(params)
	if err != nil {
		result = vo.Req204RespVO(0, "查询错误", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", likes)
	return
}

// 查看自己（别人）加入的帖吧
func (u *UserController) GetClub() (result *vo.RespVO) {
	params := u.Ctx.URLParams()
	if _, ok := params["userUuid"]; !ok {
		userUuid, _, _ := tools.ParseHeaderToken(u.Ctx)
		params["userUuid"] = userUuid
	}
	userClubs, count, err := u.AttrUserClubService.GetUserClubManySer(params)
	if err != nil {
		result = vo.Req204RespVO(0, "查询错误", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", userClubs)
	return
}

// 修改用户信息
func (u *UserController) PostInfo() (result *vo.RespVO) {
	var user domain.TbUser
	if err := u.Ctx.ReadJSON(&user); err != nil {
		result = vo.Req204RespVO(0, "参数错误", nil)
		return
	}
	user.Uuid, _, _ = tools.ParseHeaderToken(u.Ctx)
	if ok := u.AttrUserService.UserUpdateInfoSer(user); ok {
		result = vo.Req204RespVO(0, "修改失败！", nil)
		return
	}
	result = vo.Req200RespVO(1, "修改成功！", nil)
	return
}
