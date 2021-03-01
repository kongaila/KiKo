package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
)

type CommentRepository interface {
	CreateRepo(comment domain.TbComment) bool
	LikeRepo(s string) bool
	GetCommentManyRepo(params map[string]string) ([]domain.TbComment, int)
	GetSonCommentRepo(s string) []domain.TbComment
}

func NewCommentRepository(source *gorm.DB) CommentRepository {
	return &commentRepository{source: source}
}

type commentRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (c *commentRepository) GetSonCommentRepo(uuid string) (sonCom []domain.TbComment) {
	//c.source.Model(&[]domain.TbComment{}).Where("p_uuid = ? and type = ?", uuid, 2).Find(&sonCom)
	c.source.Table("tb_comment c").Select("c.*, u.*").Joins("left join tb_user u on c.user_uuid = u.uuid").Where("c.p_uuid = ? and c.type = ?", uuid, 2).Order("c.created_at desc").Find(&sonCom)
	return
}

func (c *commentRepository) GetCommentManyRepo(params map[string]string) (comments []domain.TbComment, count int) {
	articleUuid := params["articleUuid"]
	db := c.source
	if !strings.EqualFold(articleUuid, "") {
		db = db.Where("c.article_uuid = ? ", articleUuid)
	}
	userUuid := params["userUuid"]
	_, ok := params["flag"]
	if !strings.EqualFold(userUuid, "") && ok {
		db = db.Where("c.user_uuid = ? ", userUuid)
	}
	if _, ok := params["type"]; ok {
		db = db.Where("c.type = ? ", 1)
	}
	db = db.Where("c.status != ? ", 2)
	page, e1 := strconv.Atoi(params["page"])
	limit, e2 := strconv.Atoi(params["limit"])
	if e1 != nil || e2 != nil {
		return
	}
	db = db.Table("tb_comment c").Select("c.*, u.*")
	db.Count(&count)
	db.Limit(limit).Offset((page - 1) * limit).Joins("left join tb_user u on c.user_uuid = u.uuid").Order("c.created_at desc").Find(&comments)
	return
}

func (c *commentRepository) LikeRepo(uuid string) bool {
	if err := c.source.Model(&domain.TbComment{}).Where("uuid = ? ", uuid).Update("like", gorm.Expr("`like` + ?", 1)).Error; err != nil {
		return false
	}
	return true
}

func (c *commentRepository) CreateRepo(comment domain.TbComment) bool {
	//if err := c.source.Table("tb_comment").Select("id","uuid","p_id","p_uuid","type","user_uuid","article_uuid","content","created_at","status","like","article_title","report_msg").Create(&comment).Error; err != nil {
	//	return false
	//}
	if err := c.source.Exec("INSERT INTO `tb_comment` (`uuid`,`p_id`,`p_uuid`,`type`,`user_uuid`,`article_uuid`,`content`,`created_at`,`status`,`article_title`)VALUES(?,?,?,?,?,?,?,?,?,?)", comment.Uuid, comment.PId, comment.PUuid, comment.Type, comment.UserUuid, comment.ArticleUuid, comment.Content, comment.CreatedAt, comment.Status, comment.ArticleTitle).Error; err != nil {
		return false
	}
	return true
}
