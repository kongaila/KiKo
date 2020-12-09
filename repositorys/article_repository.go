package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
)

type ArticleRepository interface {
	GetArticleMany(map[string]string) ([]domain.TbArticle, int, error)
}

func NewArticleRepository(source *gorm.DB) ArticleRepository {
	return &articleRepository{source: source}
}

type articleRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (a *articleRepository) GetArticleMany(params map[string]string) (articles []domain.TbArticle, count int, err error) {
	uuid := params["uuid"]
	db := a.source
	if !strings.EqualFold(uuid, "") {
		db = db.Where("club_uuid = ? ", uuid)
	}
	// 获取总条数
	db.Count(&count)
	page, _ := strconv.Atoi(params["page"])
	limit, _ := strconv.Atoi(params["limit"])
	if page != 0 && limit != 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}
	// TODO 添加可能的随机
	db.Order("open_num desc ").Find(&articles)

	return

}
