package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

// 定义一个查询对象
type Query func(user domain.TbUser) bool

type UserRepository interface {
}

func NewUserRepository(source *gorm.DB) UserRepository {
	return &userRepository{source: source}
}

type userRepository struct {
	source *gorm.DB
	mu     sync.RWMutex
}
