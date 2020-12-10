package service

import (
	"QiqiLike/repositorys"
)

type AdminService interface {
}

func NewAdminService(repo repositorys.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

type adminService struct {
	repo repositorys.AdminRepository
}
