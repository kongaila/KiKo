package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/kataras/iris/v12"
	"time"
)

type ClubController struct {
	AttrLoginService    service.LoginService
	AttrClubService     service.ClubService
	AttrUserClubService service.UserClubService
	AttrArticleService  service.ArticleService
	AttrUserService     service.UserService
	Ctx                 iris.Context
}

// 创建贴吧
func (c *ClubController) PostCreate() (result *vo.RespVO) {
	club := domain.TbClub{}
	if err := c.Ctx.ReadJSON(&club); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	if club.MasterUuid, club.MasterName, _ = tools.ParseHeaderToken(c.Ctx); club.MasterUuid == "" || club.MasterName == "" {
		result = vo.Req204RespVO(1, "登录信息有误", nil)
		return
	}
	if ok := c.AttrClubService.Create(&club); !ok {
		result = vo.Req204RespVO(1, "添加失败", nil)
		return
	}
	userClub := domain.TbUserClub{
		Uuid:        tools.GenerateUUID(),
		ClubUuid:    club.Uuid,
		UserUuid:    club.MasterUuid,
		Identity:    1, // 吧主
		CreatedAt:   time.Now(),
		ClubName:    club.Name,
		ClubHeadImg: club.HeadImg,
	}
	c.AttrUserClubService.Create(&userClub)
	result = vo.Req200RespVO(1, "添加成功", club.Uuid)
	return
}

// 获得贴吧列表
func (c *ClubController) GetMany() (result *vo.RespVO) {
	params := c.Ctx.URLParams()
	data, count, err := c.AttrClubService.GetClubManySer(params)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", data)
	return
}

// 获得一个贴吧详情
func (c *ClubController) GetBy(uuid string) (result *vo.RespVO) {
	var club domain.TbClub
	var err error
	club, err = c.AttrClubService.GetClubDetail(uuid)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(1, "查询成功", club)
	return
}

// 获得一个贴吧的全部帖子(分页) 传uuid则查询该贴吧的帖子， 否则查询全部帖子
func (c *ClubController) GetArticleMany() (result *vo.RespVO) {
	params := c.Ctx.URLParams()
	articles, count, err := c.AttrArticleService.GetArticleManySer(params)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", articles)
	return
}

// 获得一个贴吧的全部用户
func (c *ClubController) GetUserMany() (result *vo.RespVO) {
	params := c.Ctx.URLParams()
	users, count, err := c.AttrUserService.GetUserManySer(params)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", users)
	return
}

// 加入贴吧
func (c *ClubController) PostAddBy(clubUuid string) (result *vo.RespVO) {
	userUuid, _, _ := tools.ParseHeaderToken(c.Ctx)
	userClub, _ := c.AttrClubService.GetClubDetail(clubUuid)
	if ok := c.AttrUserClubService.GetIsJoinClub(userUuid, clubUuid); ok {
		result = vo.Req204RespVO(0, "您已经加入该贴吧", nil)
		return
	}

	userClubNew := domain.TbUserClub{
		Uuid:      tools.GenerateUUID(),
		ClubUuid:  clubUuid,
		UserUuid:  userUuid,
		Identity:  0, // 成员
		CreatedAt: time.Now(),
		ClubName:  userClub.Name,
	}
	if ok := c.AttrUserClubService.Create(&userClubNew); !ok {
		result = vo.Req500RespVO(0, "加入失败", nil)
		return
	}
	result = vo.Req200RespVO(0, "加入成功", nil)
	return
}

// 退出贴吧
func (c *ClubController) PostExitBy(clubUuid string) (result *vo.RespVO) {
	userUuid, _, _ := tools.ParseHeaderToken(c.Ctx)
	if err := c.AttrUserClubService.DeleteByUserAndClubUuid(userUuid, clubUuid); err != nil {
		result = vo.Req500RespVO(0, "退出失败", nil)
		return
	}
	result = vo.Req200RespVO(0, "退出成功", nil)
	return
}
