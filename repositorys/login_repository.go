package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type LoginQuery func(user domain.TbUser) bool

type LoginRepository interface {
	InsertUser(user *domain.TbUser) (string, error)
	CheckUser(userName string) bool
	Detail(user *domain.TbUser) (ok bool)
}

func NewLoginRepository(source *gorm.DB) LoginRepository {
	return &loginRepository{source: source}
}

type loginRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (l *loginRepository) Detail(user *domain.TbUser) (ok bool) {
	var i int
	l.source.Table("tb_user").Where("user_name = ? and pass = ? ", user.UserName, user.Pass).Count(&i)
	l.source.Where("user_name = ? and pass = ? ", user.UserName, user.Pass).First(&user)
	if i == 1 {
		ok = true
		return
	}
	ok = false
	return
}

func (l *loginRepository) CheckUser(userName string) bool {
	var i int
	l.source.Table("tb_user").Where("user_name = ?", userName).Count(&i)
	if i == 0 {
		return true
	}
	return false
}

func (l *loginRepository) InsertUser(user *domain.TbUser) (string, error) {
	db := l.source.Create(user)
	return user.Uuid, db.Error
}
