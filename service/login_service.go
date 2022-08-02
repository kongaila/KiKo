package service

import (
	cs "KiKo/constants"
	"KiKo/datamodels/domain"
	"KiKo/repositorys"
	"KiKo/tools"
	"KiKo/tools/redis"
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
	token, _ = tools.GetJWTString(user.UserName, user.Uuid)
	err := redis.SetEx(token, token, time.Minute.Seconds()*30)
	if err != nil {
		return
	}
	return
}

func (l *loginService) Create(user *domain.TbUser) (string, error) {
	// 校验账号是否重复
	if ok := l.repo.CheckUser(user.UserName); !ok {
		return "", errors.New("账号名已存在")
	}
	user.Pass = tools.GeneratePass(user.Pass)
	user.CreatedAt = time.Now()
	user.Creater = user.UserName
	user.Uuid = tools.GenerateUUID()
	uuid, err := l.repo.InsertUser(user)
	if err != nil {
		log.Println(cs.ERROR, "创建用户错误： ", err)
		return "", err
	}
	return uuid, nil
}
