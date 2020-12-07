package middleware

import "github.com/kataras/iris/v12"

func RegisterMiddleware(app *iris.Application) {
	app.OnErrorCode(iris.StatusInternalServerError, error500Handler)
	app.OnErrorCode(iris.StatusBadRequest, error500Handler)
	app.Use(message, checkToken, customRecover)
	app.Done(after)
}
