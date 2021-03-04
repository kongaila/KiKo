package service

import (
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
	"KiKo/tools"
	"KiKo/tools/redis"
	"time"
)

type AdminService interface {
	AdminDetail(admin *domain.TbAdmin) (token string, ok bool)
}

func NewAdminService(repo repositorys.AdminRepository) AdminService {
	return &adminService{repo: repo}
}

type adminService struct {
	repo repositorys.AdminRepository
}

func (a adminService) AdminDetail(admin *domain.TbAdmin) (token string, ok bool) {
	admin.Pass = tools.GeneratePass(admin.Pass)
	if ok = a.repo.AdminDetail(admin); !ok {
		return
	}
	token, _ = tools.GetJWTString(admin.UserName, admin.Uuid)
	err := redis.SetEx(token, token, time.Minute.Seconds()*60*6)
	if err != nil {
		return
	}
	return
}
