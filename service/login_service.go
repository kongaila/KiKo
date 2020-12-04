package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"log"
	"time"
)

type LoginService interface {
	Create(user *domain.TbUser) (string, error)
}

func NewLoginService(repo repositorys.LoginRepository) LoginService {
	return &loginService{repo: repo}
}

type loginService struct {
	repo repositorys.LoginRepository
}

func (l *loginService) Create(user *domain.TbUser) (string, error) {
	user.GeneratePass()
	user.CreateAt = time.Now()
	user.Uuid = tools.GenerateUUID()
	uuid, err := l.repo.InsertUser(user)
	if err != nil {
		log.Println("[ERROR] 创建用户错误： ", err)
		return "", err
	}
	return uuid, nil
}
