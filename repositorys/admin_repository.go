package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type AdminRepository interface {
	AdminDetail(admin *domain.TbAdmin) bool
}

func NewAdminRepository(source *gorm.DB) AdminRepository {
	return &adminRepository{source: source}
}

type adminRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (a *adminRepository) AdminDetail(admin *domain.TbAdmin) (ok bool) {
	var i int
	a.source.Table("tb_admin").Where("user_name = ? and pass = ? ", admin.UserName, admin.Pass).Count(&i)
	a.source.Where("user_name = ? and pass = ? ", admin.UserName, admin.Pass).First(&admin)
	if i == 1 {
		ok = true
		return
	}
	ok = false
	return
}
