package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type UserRepository interface {
	GetUserManyRepo(map[string]string) ([]domain.TbUser, int, error)
	GetUserDetailRepo(s string) domain.TbUser
	UserUpdateRepo(uuid, sql string, args *gorm.SqlExpr)
	UserUpdateInfoRepo(user domain.TbUser) bool
}

func NewUserRepository(source *gorm.DB) UserRepository {
	return &userRepository{source: source}
}

type userRepository struct {
	source *gorm.DB
	mu     sync.RWMutex
}

func (u *userRepository) UserUpdateInfoRepo(user domain.TbUser) bool {
	if err := u.source.Model(&domain.TbUser{}).Select("nick", "head_img", "phone").Where("uuid = ?", user.Uuid).Updates(user); err != nil {
		return false
	}
	return true
}

// 修改用户信息
func (u *userRepository) UserUpdateRepo(uuid, sql string, args *gorm.SqlExpr) {
	u.source.Table("tb_user").Where("uuid = ?", uuid).Update(sql, args)
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
	if userName, ok := params["userName"]; ok {
		db = db.Where("tu.user_name like ?", "%"+userName+"%")
	}
	if phone, ok := params["phone"]; ok {
		db = db.Where("tu.phone like ?", "%"+phone+"%")
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
