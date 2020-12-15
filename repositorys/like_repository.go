package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
)

type LikeRepository interface {
	CreateRepo(like domain.TbLike) (string, error)
	CheckLikeRepo(like domain.TbLike) bool
	GetLikeManyRepo(strings map[string]string) ([]domain.TbLike, int, error)
}

func NewLikeRepository(source *gorm.DB) LikeRepository {
	return &likeRepository{source: source}
}

type likeRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (l *likeRepository) GetLikeManyRepo(params map[string]string) (likes []domain.TbLike, count int, err error) {
	db := l.source
	userUuid := params["userUuid"]
	if !strings.EqualFold(userUuid, "") {
		db = db.Where("user_uuid = ?", userUuid)
	}
	page, e1 := strconv.Atoi(params["page"])
	limit, e2 := strconv.Atoi(params["limit"])
	if e1 != nil || e2 != nil {
		return
	}
	if err = db.Model(&domain.TbLike{}).Count(&count).Error; err != nil {
		return
	}
	err = db.Model(&domain.TbLike{}).Limit(limit).Offset((page - 1) * limit).Order("created_at desc").Find(&likes).Error
	return
}

func (l *likeRepository) CheckLikeRepo(like domain.TbLike) bool {
	var count int
	l.source.Model(&domain.TbLike{}).Where("user_uuid = ? and article_uuid = ?", like.UserUuid, like.ArticleUuid).Count(&count)
	if count > 0 {
		return false
	}
	return true
}

func (l *likeRepository) CreateRepo(like domain.TbLike) (string, error) {
	if err := l.source.Create(&like).Error; err != nil {
		return "", err
	}
	return like.Uuid, nil
}
