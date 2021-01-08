package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
)

type DictService interface {
	GetByCodeSer(code string) []domain.TbDict
}

func NewDictService(repo repositorys.DictRepository) DictService {
	return &dictService{repo: repo}
}

type dictService struct {
	repo repositorys.DictRepository
}

func (d *dictService) GetByCodeSer(code string) []domain.TbDict {
	return d.repo.GetByCodeRepo(code)
}
