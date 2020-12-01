package main

import (
	"QiqiLike/tools"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"log"
)

func myHandler(ctx iris.Context) {
	//如果解密成功，将会进入这里,获取解密了的token
	token := ctx.Values().Get("jwt").(*jwt.Token)
	//或者这样
	//userMsg :=ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	// userMsg["id"].(float64) == 1
	// userMsg["nick_name"].(string) == iris

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	//可以了解一下token的数据结构
	ctx.Writef("%s", token.Signature)
}

func main() {

	s, err := tools.GetJWTString("kongkong", "wroqjhfafakfaffasf")
	if err != nil {
		log.Println("错误： ", err)
	}
	fmt.Println(s)

	token, err := tools.ParseToken(s)
	if err != nil {
		log.Println("错误： ", err)
	}
	fmt.Println(token)

	//app := iris.New()
	//
	//app.Use(middleware.JwtMiddle())
	//app.Get("/ping", myHandler)
	//app.Run(iris.Addr("localhost:3001"))
}
