package controllers

import (
	"KiKo/conf"
	"KiKo/datamodels/vo"
	"KiKo/tools"
	"github.com/kataras/iris/v12"
	"path"
	"path/filepath"
	"time"
)

// 文件上传控制器
type UploadController struct {
	Ctx iris.Context
}

const maxSize = 5 << 20 // 5MB

func (u *UploadController) Post() (result *vo.RespVO) {
	f, fh, err := u.Ctx.FormFile("uploadFile")
	if err != nil {
		result = vo.Req204RespVO(0, "上传文件错误", nil)
		return
	}
	defer f.Close()
	fileName := time.Now().Format("20060102") + "_" + tools.GenerateUserNick() + path.Ext(fh.Filename)
	_, err = u.Ctx.SaveFormFile(fh, filepath.Join(conf.Viper.GetString("upload.image"), fileName))
	if err != nil {
		result = vo.Req204RespVO(0, "上传文件错误", nil)
		return
	}
	result = vo.Req200RespVO(0, "上传文件成功", fileName)
	return
}
