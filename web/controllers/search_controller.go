package controllers

import (
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"github.com/kataras/iris/v12"
)

type SearchController struct {
	AttrUserService    service.UserService
	AttrClubService    service.ClubService
	AttrArticleService service.ArticleService
	Ctx                iris.Context
}

// 搜索
func (s *SearchController) GetSearch() (result *vo.RespVO) {
	var (
		r     interface{}
		count int
	)
	t, _ := s.Ctx.URLParamInt("type")
	switch t {
	// 用户
	case 1:
		r, count, _ = s.AttrUserService.GetUserManySer(s.Ctx.URLParams())
	// 帖子
	case 2:
		r, count, _ = s.AttrArticleService.GetArticleManySer(s.Ctx.URLParams())
	// 贴吧
	case 3:
		r, count, _ = s.AttrClubService.GetClubManySer(s.Ctx.URLParams())
	default:
		// 默认查询用户
		r, count, _ = s.AttrUserService.GetUserManySer(s.Ctx.URLParams())
	}

	result = vo.Req200RespVO(count, "查询成功", r)
	return
}
