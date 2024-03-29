package service

import (
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
)

type UserClubService interface {
	Create(club *domain.TbUserClub) (ok bool)
	GetUserClubManySer(strings map[string]string) ([]domain.TbUserClub, int, error)
	SelectByUuidSer(uuid string) domain.TbUserClub
	DeleteByUserAndClubUuid(userUuid string, clubUuid string) error
	GetIsJoinClub(userUuid string, clubUuid string) bool
}

func NewUserClubService(repo repositorys.UserClubRepository) UserClubService {
	return &userClubService{repo: repo}
}

type userClubService struct {
	repo repositorys.UserClubRepository
}

func (c *userClubService) GetIsJoinClub(userUuid string, clubUuid string) bool {
	if c.repo.GetIsJoinClubResp(userUuid, clubUuid) > 0 {
		return true
	} else {
		return false
	}
}

func (c *userClubService) DeleteByUserAndClubUuid(userUuid string, clubUuid string) error {
	return c.repo.DeleteByUserAndClubUuidRepo(userUuid, clubUuid)
}

func (c *userClubService) SelectByUuidSer(uuid string) domain.TbUserClub {
	return c.repo.SelectByUuidRepo(uuid)
}

func (c *userClubService) GetUserClubManySer(params map[string]string) ([]domain.TbUserClub, int, error) {
	return c.repo.GetUserClubManyRepo(params)
}

func (c *userClubService) Create(club *domain.TbUserClub) (ok bool) {
	ok = c.repo.Create(club)
	return
}
