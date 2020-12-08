package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type UserClubRepository interface {
	Create(club *domain.TbUserClub) bool
}

func NewUserClubRepository(source *gorm.DB) UserClubRepository {
	return &userClubRepository{source: source}
}

type userClubRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (c *userClubRepository) Create(club *domain.TbUserClub) bool {
	if c.source.Create(&club).Error != nil {
		return false
	}
	return true
}
