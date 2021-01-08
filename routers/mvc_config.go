package routers

import (
	"QiqiLike/web/controllers"
	"github.com/kataras/iris/v12/mvc"
)

// 登录模块
func login(app *mvc.Application) {
	app.Register(getLoginService())
	app.Handle(new(controllers.LoginController))
}

// 贴吧模块
func club(app *mvc.Application) {
	app.Register(getLoginService(), getClubService(), getUserClubService(), getArticleService(), getUserService())
	app.Handle(new(controllers.ClubController))
}

// 帖子模块
func article(app *mvc.Application) {
	app.Register(getLoginService(), getClubService(), getUserClubService(), getArticleService(), getUserService())
	app.Handle(new(controllers.ArticleController))
}

// 用户模块
func user(app *mvc.Application) {
	app.Register(getUserService(), getLikeService(), getArticleService(), getClubService(), getSearchService(), getCommentService(), getUserClubService())
	app.Handle(new(controllers.UserController))
}

// 评论模块
func comment(app *mvc.Application) {
	app.Register(getCommentService())
	app.Handle(new(controllers.CommentController))
}

// 收藏模块
func like(app *mvc.Application) {
	app.Register()
	app.Handle(new(controllers.LikeController))
}

// 搜索模块
func search(app *mvc.Application) {
	app.Register(getUserService(), getArticleService(), getClubService())
	app.Handle(new(controllers.SearchController))
}

// 热榜模块
func top(app *mvc.Application) {
	app.Register(getTopService())
	app.Handle(new(controllers.TopController))
}

// 管理员模块
func admin(app *mvc.Application) {
	app.Register()
	app.Handle(new(controllers.AdminController))
}

// 公告栏模块
func bulletin(app *mvc.Application) {
	app.Register(getBulletinService())
	app.Handle(new(controllers.BulletinController))
}

// 字典模块
func dict(app *mvc.Application) {
	app.Register(getDictService())
	app.Handle(new(controllers.DictController))
}
