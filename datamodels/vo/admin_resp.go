package vo

import cs "KiKo/constants"

// 统一Admin响应对象
type AdminRespVO struct {
	Msg   string      `json:"msg"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

// 统一调用此方法来返回响应实体
func NewAdminRespVORespVO(code, count int, msg string, data interface{}) *AdminRespVO {
	return &AdminRespVO{
		Code:  code,
		Count: count,
		Msg:   msg,
		Data:  data,
	}
}

// 成功返回响应实体
func ReqSuccessAdminRespVO(count int, msg string, data interface{}) *AdminRespVO {
	return NewAdminRespVORespVO(cs.StatusAdminSuccess, count, msg, data)
}

// 404返回响应实体
func Req404AdminRespVO(count int, msg string, data interface{}) *AdminRespVO {
	return NewAdminRespVORespVO(cs.StatusNotFound, count, msg, data)
}

// 500返回响应实体
func Req500AdminRespVO(count int, msg string, data interface{}) *AdminRespVO {
	return NewAdminRespVORespVO(cs.StatusInternalServerError, count, msg, data)
}

// err返回响应实体
func ReqFailedAdminRespVO(count int, msg string, data interface{}) *AdminRespVO {
	return NewAdminRespVORespVO(cs.StatusAdminFailed, count, msg, data)
}
