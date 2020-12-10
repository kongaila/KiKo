package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type SearchRepository interface {
}

func NewSearchRepository(source *gorm.DB) SearchRepository {
	return &searchRepository{source: source}
}

type searchRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
