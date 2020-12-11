package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/kataras/iris/v12"
)

type CommentController struct {
	AttrCommentService service.CommentService
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
	result = vo.Req200RespVO(1, "发布成功", uuid)
	return
}

// 点赞
func (c *CommentController) PostLikeBy(uuid string) (result *vo.RespVO) {
	if ok := c.AttrCommentService.LikeSer(uuid); !ok {
		result = vo.Req200RespVO(0, "点赞失败", nil)
		return
	}
	result = vo.Req200RespVO(1, "点赞成功", nil)
	return
}
