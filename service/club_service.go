package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"time"
)

type ClubService interface {
	GetClubMany(params map[string]string) (data []domain.TbClub, count int, err error)
	Create(club *domain.TbClub) (ok bool)
}

func NewClubService(repo repositorys.ClubRepository) ClubService {
	return &clubService{repo: repo}
}

type clubService struct {
	repo repositorys.ClubRepository
}

func (c *clubService) Create(club *domain.TbClub) (ok bool) {
	// 设置各种属性
	club.Uuid = tools.GenerateUUID()
	club.CreateAt = time.Now()
	club.Creater = club.MasterUuid
	club.ArticleNum = 0
	club.MemberNum = 1
	club.AttentionNum = 1
	club.Status = 0 // 待审核
	ok = c.repo.Create(club)
	return
}

func (c *clubService) GetClubMany(params map[string]string) (data []domain.TbClub, count int, err error) {
	return c.repo.GetClubMany(params)
}
