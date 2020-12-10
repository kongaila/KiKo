package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type UserService interface {
	GetUserManySer(map[string]string) ([]domain.TbUser, int, error)
}

func NewUserService(repo repositorys.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositorys.UserRepository
}

func (u *userService) GetUserManySer(params map[string]string) ([]domain.TbUser, int, error) {
	return u.repo.GetUserManyRepo(params)
}
