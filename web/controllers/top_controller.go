package controllers

import (
	"KiKo/datamodels/domain"
	"KiKo/datamodels/vo"
	"KiKo/service"
	"KiKo/tools/redis"
	"encoding/json"
)

// 其实大多数时候 我都很爱你
type TopController struct {
	AttrTopService service.TopService
}

// 获得热榜信息
func (t *TopController) Get() (result *vo.RespVO) {
	tops, _ := redis.GetString("top")
	var data domain.TopSlice
	_ = json.Unmarshal([]byte(tops), &data)
	result = vo.Req200RespVO(10, "查询成功", data)
	return
}

// 查看历史热榜
func (t *TopController) GetBy(topNum int) (result *vo.RespVO) {
	data := t.AttrTopService.GetHistoryTopSer(topNum)
	result = vo.Req200RespVO(10, "查询成功", data)
	return
}
