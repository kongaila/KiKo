package service

import (
	cs "QiqiLike/constants"
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"errors"
	"log"
	"time"
)

type LoginService interface {
	Create(user *domain.TbUser) (string, error)
	UserDetail(user *domain.TbUser) (token string, ok bool)
}

func NewLoginService(repo repositorys.LoginRepository) LoginService {
	return &loginService{repo: repo}
}

type loginService struct {
	repo repositorys.LoginRepository
}

func (l *loginService) UserDetail(user *domain.TbUser) (token string, ok bool) {
	user.Pass = tools.GeneratePass(user.Pass)
	if ok = l.repo.Detail(user); !ok {
		return
	}
	// TODO 进行获取token操作
	token, _ = tools.GetJWTString(user.UserName, user.Uuid)

	return
}

func (l *loginService) Create(user *domain.TbUser) (string, error) {
	// 校验账号是否重复
	if ok := l.repo.CheckUser(user.UserName); !ok {
		return "", errors.New("账号名已存在")
	}
	tools.GeneratePass(user.Pass)
	user.CreateAt = time.Now()
	user.Creater = user.UserName
	user.Uuid = tools.GenerateUUID()
	uuid, err := l.repo.InsertUser(user)
	if err != nil {
		log.Println(cs.ERROR, "创建用户错误： ", err)
		return "", err
	}
	return uuid, nil
}
