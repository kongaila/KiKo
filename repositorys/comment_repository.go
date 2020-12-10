package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type CommentRepository interface {
}

func NewCommentRepository(source *gorm.DB) CommentRepository {
	return &commentRepository{source: source}
}

type commentRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
