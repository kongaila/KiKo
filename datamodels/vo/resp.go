package vo

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
		Msg:   msg,
		Code:  code,
		Data:  data,
		Count: count,
	}
}
