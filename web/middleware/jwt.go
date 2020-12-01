package middleware

import (
	cs "QiqiLike/constants"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

//解释：
//jwtmiddleware.New是配置中间件的错误返回，是否为调试模式，机密秘钥，加密模式等
//app.Use(jwtHandler.Serve) 是把中间件注册到处理程序中
//注册次中间件的路由，中间件每一次都会去获取header头Authorization字段用户判断
func JwtMiddle() func(ctx iris.Context) {
	//jwt中间件
	return jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(cs.Salt), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)
	}).Serve
}
