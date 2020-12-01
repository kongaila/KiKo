package vo

import cs "QiqiLike/constants"

// 统一响应对象
type respVO struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

// 统一调用此方法来返回响应实体
func NewRespVO(code, count int, msg string, data interface{}) *respVO {
	return &respVO{
		Code:  code,
		Count: count,
		Msg:   msg,
		Data:  data,
	}
}

// 成功返回响应实体
func Req200RespVO(count int, data interface{}) *respVO {
	return NewRespVO(cs.StatusOK, count, cs.StatusText[cs.StatusOK], data)
}

// 404返回响应实体
func Req404RespVO(count int, data interface{}) *respVO {
	return NewRespVO(cs.StatusNotFound, count, cs.StatusText[cs.StatusNotFound], data)
}

// 500返回响应实体
func Req500RespVO(count int, data interface{}) *respVO {
	return NewRespVO(cs.StatusInternalServerError, count, cs.StatusText[cs.StatusInternalServerError], data)
}

// 204返回响应实体
func Req204RespVO(count int, data interface{}) *respVO {
	return NewRespVO(cs.StatusNoContent, count, cs.StatusText[cs.StatusNoContent], data)
}
