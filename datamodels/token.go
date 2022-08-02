package datamodels

// Token结构体
type Token struct {
	Uuid     string
	UserName string
	// 签发人
	Iss string
	// 签发时间
	Iat int64
	// 过期时间
	Exp int64
}
