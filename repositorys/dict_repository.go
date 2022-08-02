package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type DictRepository interface {
	GetByCodeRepo(code string) []domain.TbDict
}

func NewDictRepository(source *gorm.DB) DictRepository {
	return &dictRepository{source: source}
}

type dictRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (d *dictRepository) GetByCodeRepo(code string) (dicts []domain.TbDict) {
	var id int32
	_ = d.source.Raw("SELECT id FROM tb_dict WHERE code = ?", code).Row().Scan(&id)
	d.source.Where("pid = ?", id).Find(&dicts)
	return
}
