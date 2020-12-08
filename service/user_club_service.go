package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type UserClubService interface {
	Create(club *domain.TbUserClub) (ok bool)
}

func NewUserClubService(repo repositorys.UserClubRepository) UserClubService {
	return &userClubService{repo: repo}
}

type userClubService struct {
	repo repositorys.UserClubRepository
}

func (c *userClubService) Create(club *domain.TbUserClub) (ok bool) {
	ok = c.repo.Create(club)
	return
}
