package middleware

import (
	cs "QiqiLike/constants"
	"github.com/kataras/iris/v12"
	"log"
)

// 最后执行（目前没有用到）
func after(ctx iris.Context) {
	log.Println(cs.INFO, "走完了!")
}
