package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
)

type UserClubRepository interface {
	Create(club *domain.TbUserClub) bool
	GetUserClubManyRepo(strings map[string]string) ([]domain.TbUserClub, int, error)
	SelectByUuidRepo(uuid string) domain.TbUserClub
	DeleteByUserAndClubUuidRepo(string, string) error
}

func NewUserClubRepository(source *gorm.DB) UserClubRepository {
	return &userClubRepository{source: source}
}

type userClubRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (c *userClubRepository) DeleteByUserAndClubUuidRepo(userUuid string, clubUuid string) error {
	return c.source.Where("user_uuid = ? and club_uuid = ?", userUuid, clubUuid).Delete(&domain.TbUserClub{}).Error
}

func (c *userClubRepository) SelectByUuidRepo(uuid string) (tuc domain.TbUserClub) {
	c.source.Model(&domain.TbUserClub{}).Where("uuid = ?", uuid).First(&tuc)
	return
}

func (c *userClubRepository) GetUserClubManyRepo(params map[string]string) (userClubs []domain.TbUserClub, count int, err error) {
	db := c.source
	userUuid := params["userUuid"]
	if !strings.EqualFold(userUuid, "") {
		db = db.Where("user_uuid = ?", userUuid)
	}
	page, e1 := strconv.Atoi(params["page"])
	limit, e2 := strconv.Atoi(params["limit"])
	if e1 != nil || e2 != nil {
		return
	}
	if err = db.Model(&domain.TbUserClub{}).Count(&count).Error; err != nil {
		return
	}
	err = db.Model(&domain.TbUserClub{}).Limit(limit).Offset((page - 1) * limit).Order("created_at desc").Find(&userClubs).Error
	return
}

func (c *userClubRepository) Create(club *domain.TbUserClub) bool {
	if c.source.Create(&club).Error != nil {
		return false
	}
	return true
}
