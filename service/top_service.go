package service

import (
	"QiqiLike/repositorys"
)

type TopService interface {
}

func NewTopService(repo repositorys.TopRepository) TopService {
	return &topService{repo: repo}
}

type topService struct {
	repo repositorys.TopRepository
}
