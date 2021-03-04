package tools

import (
	"KiKo/conf"
	cs "KiKo/constants"
	"KiKo/datamodels"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
)

// 获取有效期为永久的jwt字符串（100年）
func GetJWTString(userName, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": userName,
		"uuid":     uuid,
		"iss":      conf.Viper.GetString("title"),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add((24 * 365 * 100) * time.Hour * time.Duration(1)).Unix(), // 100年
	})
	// 把token已约定的加密方式和加密秘钥加密
	tokenString, err := token.SignedString([]byte(cs.Salt))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析jwt字符串
func ParseToken(token string) (*datamodels.Token, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(cs.Salt), nil
	})
	if err != nil {
		return nil, err
	}
	t := datamodels.Token{}
	err = mapstructure.Decode(claim.Claims.(jwt.MapClaims), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// 封装好的方法直接返回uuid和用户名
func ParseTokenUuidAndUserName(token string) (string, string, error) {
	t, err := ParseToken(token)
	if err != nil {
		return "", "", err
	}
	return t.Uuid, t.UserName, nil
}

// 从请求头中拿到token
func GetHeaderToken(ctx iris.Context) string {
	token := ctx.GetHeader("Authorization")
	return strings.Replace(token, "Bearer ", "", 1)
}

// 封装好的方法从请求头中直接返回uuid和username
func ParseHeaderToken(ctx iris.Context) (userUuid string, userName string, err error) {
	return ParseTokenUuidAndUserName(GetHeaderToken(ctx))
}
