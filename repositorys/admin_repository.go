package repositorys

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type AdminRepository interface {
}

func NewAdminRepository(source *gorm.DB) AdminRepository {
	return &adminRepository{source: source}
}

type adminRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}
