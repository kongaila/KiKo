package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"errors"
	"time"
)

type LikeService interface {
	CreateLike(like domain.TbLike) (string, error)
}

func NewLikeService(repo repositorys.LikeRepository, ArticleRepo repositorys.ArticleRepository) LikeService {
	return &likeService{
		repo:        repo,
		ArticleRepo: ArticleRepo,
	}
}

type likeService struct {
	repo        repositorys.LikeRepository
	ArticleRepo repositorys.ArticleRepository
}

func (l *likeService) CreateLike(like domain.TbLike) (string, error) {
	if ok := l.ArticleRepo.CheckArticle(like.ArticleUuid); !ok {
		return "", errors.New("要收藏的文章不存在")
	}
	if ok := l.repo.CheckLikeRepo(like); !ok {
		return "", errors.New("您已经收藏， 不可重复收藏")
	}
	like.Uuid = tools.GenerateUUID()
	like.CreatedAt = time.Now()
	return l.repo.CreateRepo(like)
}
