package cs

// 定义状态码以及对应消息
const (
	StatusOK                  = 200
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusGatewayTimeout      = 504
	NotToken                  = 666
)

var StatusText = map[int]string{
	StatusOK:                  "成功",
	StatusNoContent:           "响应无内容",
	StatusBadRequest:          "错误请求",
	StatusNotFound:            "访问资源不存在",
	StatusInternalServerError: "服务器异常",
	StatusBadGateway:          "过多的请求",
	StatusGatewayTimeout:      "请求超时",
	NotToken:                  "没有token值",
}
