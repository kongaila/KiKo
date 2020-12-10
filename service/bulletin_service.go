package service

import (
	"QiqiLike/repositorys"
)

type BulletinService interface {
}

func NewBulletinService(repo repositorys.BulletinRepository) BulletinService {
	return &bulletinService{repo: repo}
}

type bulletinService struct {
	repo repositorys.AdminRepository
}
