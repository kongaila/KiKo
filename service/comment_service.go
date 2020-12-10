package service

import (
	"QiqiLike/repositorys"
)

type CommentService interface {
}

func NewCommentService(repo repositorys.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

type commentService struct {
	repo repositorys.CommentRepository
}
