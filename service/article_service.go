package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type ArticleService interface {
	GetArticleMany(map[string]string) ([]domain.TbArticle, int, error)
}

func NewArticleService(repo repositorys.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

type articleService struct {
	repo repositorys.ArticleRepository
}

// 获得全部帖子，
func (a *articleService) GetArticleMany(params map[string]string) ([]domain.TbArticle, int, error) {
	return a.repo.GetArticleMany(params)
}
