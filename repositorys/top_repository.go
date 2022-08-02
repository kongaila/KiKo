package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"sync"
)

type TopRepository interface {
	GetHistoryTopRepo(i int) []domain.TbTop
}

func NewTopRepository(source *gorm.DB) TopRepository {
	return &topRepository{source: source}
}

type topRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (t *topRepository) GetHistoryTopRepo(i int) (data []domain.TbTop) {
	t.source.Table("tb_top").Where("num = ?", i).Limit(10).Find(&data)
	return
}
