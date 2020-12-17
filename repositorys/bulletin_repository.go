package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type BulletinRepository interface {
	GetBulletinRepo() domain.TbBulletin
}

func NewBulletinRepository(source *gorm.DB) BulletinRepository {
	return &bulletinRepository{source: source}
}

type bulletinRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (b *bulletinRepository) GetBulletinRepo() (data domain.TbBulletin) {
	b.source.Model(&domain.TbBulletin{}).Last(&data)
	return
}
