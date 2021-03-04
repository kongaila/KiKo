package service

import (
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
	"time"
)

type BulletinService interface {
	GetBulletinSer() domain.TbBulletin
	Create(bulletin domain.TbBulletin) bool
	GetManySer(map[string]string) (int, []domain.TbBulletin)
}

func NewBulletinService(repo repositorys.BulletinRepository) BulletinService {
	return &bulletinService{repo: repo}
}

type bulletinService struct {
	repo repositorys.BulletinRepository
}

func (b *bulletinService) GetManySer(params map[string]string) (int, []domain.TbBulletin) {
	return b.repo.GetManyRepo(params)
}

func (b *bulletinService) Create(bulletin domain.TbBulletin) bool {
	bulletin.CreatedAt = time.Now()
	return b.repo.Create(bulletin)
}

func (b *bulletinService) GetBulletinSer() domain.TbBulletin {
	return b.repo.GetBulletinRepo()
}
