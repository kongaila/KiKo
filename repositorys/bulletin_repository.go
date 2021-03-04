package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type BulletinRepository interface {
	GetBulletinRepo() domain.TbBulletin
	Create(bulletin domain.TbBulletin) bool
	GetManyRepo(params map[string]string) (int, []domain.TbBulletin)
}

func NewBulletinRepository(source *gorm.DB) BulletinRepository {
	return &bulletinRepository{source: source}
}

type bulletinRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (b *bulletinRepository) GetManyRepo(params map[string]string) (count int, bs []domain.TbBulletin) {
	page, _ := strconv.Atoi(params["page"])
	limit, _ := strconv.Atoi(params["limit"])
	db := b.source
	db.Table("tb_bulletin a").Count(&count)
	db = db.Limit(limit).Offset((page - 1) * limit)
	db.Table("tb_bulletin").Order("created_at desc").Find(&bs)
	return
}

func (b *bulletinRepository) Create(bulletin domain.TbBulletin) bool {
	err := b.source.Create(&bulletin).Error
	if err != nil {
		return false
	}
	return true
}

func (b *bulletinRepository) GetBulletinRepo() (data domain.TbBulletin) {
	b.source.Model(&domain.TbBulletin{}).Last(&data)
	return
}
