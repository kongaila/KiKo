package vo

import cs "QiqiLike/constants"

// 统一响应对象
type RespVO struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

// 统一调用此方法来返回响应实体
func NewRespVO(code, count int, msg string, data interface{}) *RespVO {
	return &RespVO{
		Code:  code,
		Count: count,
		Msg:   msg,
		Data:  data,
	}
}

// 成功返回响应实体
func Req200RespVO(count int, msg string, data interface{}) *RespVO {
	return NewRespVO(cs.StatusOK, count, msg, data)
}

// 404返回响应实体
func Req404RespVO(count int, msg string, data interface{}) *RespVO {
	return NewRespVO(cs.StatusNotFound, count, msg, data)
}

// 500返回响应实体
func Req500RespVO(count int, msg string, data interface{}) *RespVO {
	return NewRespVO(cs.StatusInternalServerError, count, msg, data)
}

// 204返回响应实体
func Req204RespVO(count int, msg string, data interface{}) *RespVO {
	return NewRespVO(cs.StatusNoContent, count, msg, data)
}

// 204返回响应实体
func ReqNotTokenRespVO(count int, msg string, data interface{}) *RespVO {
	return NewRespVO(cs.NotToken, count, msg, data)
}
