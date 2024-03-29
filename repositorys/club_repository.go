package repositorys

import (
	"KiKo/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"sync"
)

type ClubRepository interface {
	Create(club *domain.TbClub) bool
	GetClubManyRepo(params map[string]string) ([]domain.TbClub, int, error)
	GetClubDetail(s string) (domain.TbClub, error)
	UpdateClubInfoSer(s string, s2 string, s3 string) bool
}

func NewClubRepository(source *gorm.DB) ClubRepository {
	return &clubRepository{source: source}
}

type clubRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (c *clubRepository) UpdateClubInfoSer(uuid string, sql string, arg string) bool {
	if err := c.source.Table("tb_club").Where("uuid = ?", uuid).Update(sql, arg).Error; err != nil {
		return false
	}
	return true
}

func (c *clubRepository) GetClubDetail(uuid string) (domain.TbClub, error) {
	var club domain.TbClub
	db := c.source.Where("uuid = ?", uuid).First(&club)
	if db.Error != nil {
		return club, db.Error
	}
	c.source.Table("tb_article").Where("club_uuid = ?", club.Uuid).Count(&club.ArticleNum)
	c.source.Table("tb_user_club").Where("club_uuid = ?", club.Uuid).Count(&club.MemberNum)
	return club, nil
}

func (c *clubRepository) GetClubManyRepo(params map[string]string) (data []domain.TbClub, count int, err error) {
	page, err1 := strconv.Atoi(params["page"])
	limit, err2 := strconv.Atoi(params["limit"])
	if err1 != nil || err2 != nil {
		return nil, 0, err1
	}
	db := c.source
	if query, ok := params["query"]; ok {
		db = db.Where("name like ? or type_name like ?", "%"+query+"%", "%"+query+"%")
	}
	// 只查询待审核的
	if _, ok := params["check"]; ok {
		db = db.Where("status = ?", 0)
	} else {
		// 否则只查询审核通过的
		db = db.Where("status = ?", 1)
	}
	// 获取取指page，指定limit的记录
	db.Limit(limit).Offset((page - 1) * limit).Order("created_at desc").Find(&data)
	// 获取总条数
	db.Model(&domain.TbClub{}).Count(&count)
	return data, count, nil
}

func (c *clubRepository) Create(club *domain.TbClub) bool {
	if c.source.Create(&club).Error != nil {
		return false
	}
	return true
}
