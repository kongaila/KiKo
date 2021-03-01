package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	GetUserManySer(map[string]string) ([]domain.TbUser, int, error)
	GetUserDetailSer(s string) domain.TbUser
	UserUpdateSer(string, string, *gorm.SqlExpr)
}

func NewUserService(repo repositorys.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositorys.UserRepository
}

func (u *userService) UserUpdateSer(uuid, sql string, args *gorm.SqlExpr) {
	u.repo.UserUpdateRepo(uuid, sql, args)
}

func (u *userService) GetUserDetailSer(uuid string) domain.TbUser {
	return u.repo.GetUserDetailRepo(uuid)
}

func (u *userService) GetUserManySer(params map[string]string) ([]domain.TbUser, int, error) {
	return u.repo.GetUserManyRepo(params)
}
