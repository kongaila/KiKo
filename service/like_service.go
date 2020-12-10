package service

import (
	"QiqiLike/repositorys"
)

type LikeService interface {
}

func NewLikeService(repo repositorys.LikeRepository) LikeService {
	return &likeService{repo: repo}
}

type likeService struct {
	repo repositorys.LikeRepository
}
