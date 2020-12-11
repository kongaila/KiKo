package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type LikeRepository interface {
	CreateRepo(like domain.TbLike) (string, error)
	CheckLikeRepo(like domain.TbLike) bool
}

func NewLikeRepository(source *gorm.DB) LikeRepository {
	return &likeRepository{source: source}
}

type likeRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
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
