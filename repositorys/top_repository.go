package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type TopRepository interface {
}

func NewTopRepository(source *gorm.DB) TopRepository {
	return &topRepository{source: source}
}

type topRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
