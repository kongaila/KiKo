package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"time"
)

type ClubService interface {
	GetClubManySer(params map[string]string) (data []domain.TbClub, count int, err error)
	Create(club *domain.TbClub) (ok bool)
	GetClubDetail(s string) (domain.TbClub, error)
}

func NewClubService(repo repositorys.ClubRepository) ClubService {
	return &clubService{repo: repo}
}

type clubService struct {
	repo repositorys.ClubRepository
}

func (c *clubService) GetClubDetail(uuid string) (domain.TbClub, error) {
	return c.repo.GetClubDetail(uuid)
}

func (c *clubService) Create(club *domain.TbClub) (ok bool) {
	// 设置各种属性
	club.Uuid = tools.GenerateUUID()
	club.CreatedAt = time.Now()
	club.Creater = club.MasterUuid
	club.ArticleNum = 0
	club.MemberNum = 1
	club.AttentionNum = 1
	club.Status = 0 // 待审核
	ok = c.repo.Create(club)
	return
}

func (c *clubService) GetClubManySer(params map[string]string) (data []domain.TbClub, count int, err error) {
	return c.repo.GetClubManyRepo(params)
}
