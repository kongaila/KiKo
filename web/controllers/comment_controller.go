package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

type CommentController struct {
	AttrCommentService service.CommentService
	AttrUserService    service.UserService
	Ctx                iris.Context
}

// 发布评论
func (c *CommentController) PostCreate() (result *vo.RespVO) {
	comment := domain.TbComment{}
	if err := c.Ctx.ReadJSON(&comment); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	var err error
	if comment.UserUuid, _, err = tools.ParseHeaderToken(c.Ctx); err != nil {
		result = vo.Req204RespVO(0, "token异常", nil)
		return
	}
	var (
		uuid string
		ok   bool
	)
	if uuid, ok = c.AttrCommentService.CreateSer(comment); !ok {
		result = vo.Req204RespVO(0, "发布失败", nil)
		return
	}
	c.AttrUserService.UserUpdateSer(comment.UserUuid, "comment_num", gorm.Expr("comment_num + ?", 1))
	result = vo.Req200RespVO(1, "发布成功", uuid)
	return
}

// 点赞评论
func (c *CommentController) PostLikeBy(uuid string) (result *vo.RespVO) {
	if ok := c.AttrCommentService.LikeSer(uuid); !ok {
		result = vo.Req200RespVO(0, "点赞失败", nil)
		return
	}
	result = vo.Req200RespVO(1, "点赞成功", nil)
	return
}

// 查看帖子中的所有评论(只有一级)
func (c *CommentController) GetMany() (result *vo.RespVO) {
	params := c.Ctx.URLParams()
	params["type"] = "true"
	comments, count := c.AttrCommentService.GetCommentManySer(params)
	for i, v := range comments {
		sonComment := c.AttrCommentService.GetSonCommentSer(v.Uuid)
		comments[i].SonComment = sonComment
	}
	result = vo.Req200RespVO(count, "查询成功", comments)
	return
}
