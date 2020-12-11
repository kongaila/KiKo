package service

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/repositorys"
	"QiqiLike/tools"
	"time"
)

type CommentService interface {
	CreateSer(domain.TbComment) (string, bool)
	LikeSer(string) bool
}

func NewCommentService(repo repositorys.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

type commentService struct {
	repo repositorys.CommentRepository
}

func (c *commentService) LikeSer(uuid string) bool {
	return c.repo.LikeRepo(uuid)
}

func (c *commentService) CreateSer(comment domain.TbComment) (string, bool) {
	comment.Uuid = tools.GenerateUUID()
	comment.CreatedAt = time.Now()
	if ok := c.repo.CreateRepo(comment); !ok {
		return "", false
	}
	return comment.Uuid, true
}
