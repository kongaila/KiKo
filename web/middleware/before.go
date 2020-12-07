package middleware

import (
	cs "QiqiLike/constants"
	"QiqiLike/datamodels/vo"
	"QiqiLike/tools"
	"QiqiLike/tools/redis"
	"github.com/kataras/iris/v12"
	"log"
	"strings"
	"time"
)

// 不受控制的路径
var uncontrolled = []string{cs.Prefix + "/login", cs.Prefix + "/register"}

// 打印日志
func message(ctx iris.Context) {
	log.Printf(cs.INFO+"发送了请求： %s %s%s", ctx.Method(), ctx.Host(), ctx.Request().URL.String())
	ctx.Next()
}

// 校验token
func checkToken(ctx iris.Context) {
	for _, v := range uncontrolled {
		if strings.EqualFold(v, ctx.Request().URL.String()) {
			ctx.Next()
			return
		}
	}
	log.Println(cs.INFO, "校验token!")
	// 验证有没有token 没有的话提示登录
	token := tools.GetHeaderToken(ctx)
	ex, ok := redis.Exists(token)
	if !ok {
		// 不存在请求登录
		ctx.JSON(vo.ReqNotTokenRespVO(0, "请登录！", nil))
		return
	}

	// 如果剩余时间小于10分钟， 那么就将其续约至30分钟
	if ex < time.Minute.Seconds()*10 {
		_ = redis.SetEx(token, token, time.Minute.Seconds()*30)
	}

	ctx.Next()
}
