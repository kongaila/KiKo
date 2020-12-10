package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type LikeRepository interface {
}

func NewLikeRepository(source *gorm.DB) LikeRepository {
	return &likeRepository{source: source}
}

type likeRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
