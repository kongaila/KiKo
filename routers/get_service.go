package routers

import (
	repo "KiKo/repositorys"
	"KiKo/service"
)

// 获得登录Service
func getLoginService() service.LoginService {
	dao := repo.NewLoginRepository(repo.Db)
	return service.NewLoginService(dao)
}

// 获得贴吧Service
func getClubService() service.ClubService {
	dao := repo.NewClubRepository(repo.Db)
	return service.NewClubService(dao)
}

// 获得用户Service
func getUserService() service.UserService {
	dao := repo.NewUserRepository(repo.Db)
	return service.NewUserService(dao)
}

// 获得用户贴吧关联Service
func getUserClubService() service.UserClubService {
	dao := repo.NewUserClubRepository(repo.Db)
	return service.NewUserClubService(dao)
}

// 获得帖子Service
func getArticleService() service.ArticleService {
	dao := repo.NewArticleRepository(repo.Db)
	return service.NewArticleService(dao)
}

// 获得喜欢Service
func getLikeService() service.LikeService {
	dao := repo.NewLikeRepository(repo.Db)
	dao2 := repo.NewArticleRepository(repo.Db)
	return service.NewLikeService(dao, dao2)
}

// 获得评论Service
func getCommentService() service.CommentService {
	dao := repo.NewCommentRepository(repo.Db)
	return service.NewCommentService(dao)
}

// 获得搜索Service
func getSearchService() service.SearchService {
	dao := repo.NewSearchRepository(repo.Db)
	return service.NewSearchService(dao)
}

// 获得热榜Service
func getTopService() service.TopService {
	dao := repo.NewTopRepository(repo.Db)
	return service.NewTopService(dao)
}

// 获得管理员Service
func getAdminService() service.AdminService {
	dao := repo.NewAdminRepository(repo.Db)
	return service.NewAdminService(dao)
}

// 获得公告栏Service
func getBulletinService() service.BulletinService {
	dao := repo.NewBulletinRepository(repo.Db)
	return service.NewBulletinService(dao)
}

// 获得字典Service
func getDictService() service.DictService {
	dao := repo.NewDictRepository(repo.Db)
	return service.NewDictService(dao)
}
