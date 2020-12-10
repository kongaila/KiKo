package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type UserRepository interface {
	GetUserManyRepo(map[string]string) ([]domain.TbUser, int, error)
}

func NewUserRepository(source *gorm.DB) UserRepository {
	return &userRepository{source: source}
}

type userRepository struct {
	source *gorm.DB
	mu     sync.RWMutex
}

func (u *userRepository) GetUserManyRepo(params map[string]string) (users []domain.TbUser, count int, err error) {
	uuid := params["uuid"]
	db := u.source
	page, _ := strconv.Atoi(params["page"])
	limit, _ := strconv.Atoi(params["limit"])
	db = db.Table("tb_club tc").Select("tu.* ").Joins("LEFT JOIN tb_user_club tuc ON tuc.club_uuid = tc.uuid").Joins("LEFT JOIN tb_user tu ON tuc.user_uuid = tu.uuid").Where("tc.uuid  = ? ", uuid).Order("tuc.create_at desc ")
	db.Count(&count)
	if page != 0 && limit != 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}
	db.Find(&users)
	return
}
