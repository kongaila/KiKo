package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type UserRepository interface {
	GetUserManyRepo(map[string]string) ([]domain.TbUser, int, error)
	GetUserDetailRepo(s string) domain.TbUser
	UserUpdateRepo(uuid, sql string, args []interface{})
}

func NewUserRepository(source *gorm.DB) UserRepository {
	return &userRepository{source: source}
}

type userRepository struct {
	source *gorm.DB
	mu     sync.RWMutex
}

// 修改用户信息
func (u *userRepository) UserUpdateRepo(uuid, sql string, args []interface{}) {
	u.source.Where("uuid = ?", uuid).Update(sql, args)
}

func (u *userRepository) GetUserDetailRepo(uuid string) (user domain.TbUser) {
	u.source.Model(&domain.TbUser{}).Where("uuid = ?", uuid).Last(&user)
	return
}

func (u *userRepository) GetUserManyRepo(params map[string]string) (users []domain.TbUser, count int, err error) {
	db := u.source
	if clubUuid, ok := params["clubUuid"]; ok {
		db = db.Where("tuc.club_uuid = ?", clubUuid)
	}
	if query, ok := params["query"]; ok {
		db = db.Where("tu.nick like ?", "%"+query+"%")
	}
	page, _ := strconv.Atoi(params["page"])
	limit, _ := strconv.Atoi(params["limit"])

	db = db.Table("tb_user tu").Select("tu.* ").Joins("LEFT JOIN tb_user_club tuc ON tuc.user_uuid = tu.uuid").Order("tuc.created_at desc ")
	db.Count(&count)
	if page != 0 && limit != 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}
	db.Find(&users)
	return
}
