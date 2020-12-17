package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type TopService interface {
	GetHistoryTopSer(i int) []domain.TbTop
}

func NewTopService(repo repositorys.TopRepository) TopService {
	return &topService{repo: repo}
}

type topService struct {
	repo repositorys.TopRepository
}

func (t *topService) GetHistoryTopSer(i int) []domain.TbTop {
	return t.repo.GetHistoryTopRepo(i)
}
