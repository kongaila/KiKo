package routers

import (
	repo "QiqiLike/repositorys"
	"QiqiLike/service"
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
