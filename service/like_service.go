package service

import (
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
	"KiKo/tools"
	"errors"
	"time"
)

type LikeService interface {
	CreateLike(like domain.TbLike) (string, error)
	GetLikeManySer(strings map[string]string) ([]domain.TbLike, int, error)
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

func (l *likeService) GetLikeManySer(params map[string]string) ([]domain.TbLike, int, error) {
	return l.repo.GetLikeManyRepo(params)
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
