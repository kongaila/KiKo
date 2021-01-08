package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"time"
)

type ArticleService interface {
	GetArticleManySer(map[string]string) ([]domain.TbArticle, int, error)
	Create(article *domain.TbArticle) bool
	GetDetailSer(uuid string) (domain.TbArticle, error)
	ReportMsgSer(article domain.TbArticle) bool
}

func NewArticleService(repo repositorys.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

type articleService struct {
	repo repositorys.ArticleRepository
}

func (a *articleService) ReportMsgSer(article domain.TbArticle) bool {
	return a.repo.ReportMsgRepo(article)
}

func (a *articleService) GetDetailSer(uuid string) (domain.TbArticle, error) {
	return a.repo.SelectArticleDetailRepo(uuid)
}

func (a *articleService) Create(article *domain.TbArticle) bool {
	article.Uuid = tools.GenerateUUID()
	article.CreatedAt = time.Now()
	article.Flag = 1
	article.IsGood = 0
	article.Status = 0
	return a.repo.CreateRepo(article)
}

// 获得全部帖子，
func (a *articleService) GetArticleManySer(params map[string]string) ([]domain.TbArticle, int, error) {
	return a.repo.GetArticleManyRepo(params)
}
