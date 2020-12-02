package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type UserService interface {
	GetAll() []domain.TbUser
	GetByID(id int64) (domain.TbUser, bool)
	GetByUsernameAndPassword(username, userPassword string) (domain.TbUser, bool)
	DeleteByID(id int64) bool

	Update(id int64, user domain.TbUser) (domain.TbUser, error)
	UpdatePassword(id int64, newPassword string) (domain.TbUser, error)
	UpdateUsername(id int64, newUsername string) (domain.TbUser, error)

	Create(userPassword string, user domain.TbUser) (domain.TbUser, error)
}

func NewUserService(repo repositorys.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

type userService struct {
	repo repositorys.UserRepository
}

func (u *userService) GetAll() []domain.TbUser {
	panic("implement me")
}

func (u *userService) GetByID(id int64) (domain.TbUser, bool) {
	panic("implement me")
}

func (u *userService) GetByUsernameAndPassword(username, userPassword string) (domain.TbUser, bool) {
	panic("implement me")
}

func (u *userService) DeleteByID(id int64) bool {
	panic("implement me")
}

func (u *userService) Update(id int64, user domain.TbUser) (domain.TbUser, error) {
	panic("implement me")
}

func (u *userService) UpdatePassword(id int64, newPassword string) (domain.TbUser, error) {
	panic("implement me")
}

func (u *userService) UpdateUsername(id int64, newUsername string) (domain.TbUser, error) {
	panic("implement me")
}

func (u *userService) Create(userPassword string, user domain.TbUser) (domain.TbUser, error) {
	panic("implement me")
}
