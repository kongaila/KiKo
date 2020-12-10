package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type BulletinRepository interface {
}

func NewBulletinRepository(source *gorm.DB) BulletinRepository {
	return &bulletinRepository{source: source}
}

type bulletinRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
