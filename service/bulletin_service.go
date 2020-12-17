package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type BulletinService interface {
	GetBulletinSer() domain.TbBulletin
}

func NewBulletinService(repo repositorys.BulletinRepository) BulletinService {
	return &bulletinService{repo: repo}
}

type bulletinService struct {
	repo repositorys.BulletinRepository
}

func (b *bulletinService) GetBulletinSer() domain.TbBulletin {
	return b.repo.GetBulletinRepo()
}
