package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type LoginRepository interface {
	InsertUser(user *domain.TbUser) (string, error)
}

func NewLoginRepository(source *gorm.DB) *loginRepository {
	return &loginRepository{source: source}
}

type loginRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (l *loginRepository) InsertUser(user *domain.TbUser) (string, error) {
	db := l.source.Create(user)
	return user.Uuid, db.Error
}
