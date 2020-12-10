package service

import (
	"QiqiLike/repositorys"
)

type SearchService interface {
}

func NewSearchService(repo repositorys.SearchRepository) SearchService {
	return &searchService{repo: repo}
}

type searchService struct {
	repo repositorys.SearchRepository
}
