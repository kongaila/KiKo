package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type CommentRepository interface {
	CreateRepo(comment domain.TbComment) bool
	LikeRepo(s string) bool
}

func NewCommentRepository(source *gorm.DB) CommentRepository {
	return &commentRepository{source: source}
}

type commentRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (c *commentRepository) LikeRepo(uuid string) bool {
	if err := c.source.Model(&domain.TbComment{}).Where("uuid = ? ", uuid).Update("like", gorm.Expr("`like` + ?", 1)).Error; err != nil {
		return false
	}
	return true
}

func (c *commentRepository) CreateRepo(comment domain.TbComment) bool {
	if err := c.source.Model(&domain.TbComment{}).Create(&comment).Error; err != nil {
		return false
	}
	return true
}
