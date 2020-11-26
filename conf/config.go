package conf

import (
	"github.com/spf13/viper"
)

// 读取的路径, 默认dev
var path = "dev"

// viper 全局实例
var Viper *viper.Viper

// 初始化viper
func init() {
	Viper = viper.New()
	Viper.SetConfigName("config")         // 指定配置文件的文件名称 (不需要制定配置文件的扩展名 )
	Viper.AddConfigPath("./conf/" + path) // 设置配置文件和可执行二进制文件在用一个目录

	err := Viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		panic(err) // 读取配置文件失败致命错误
	}
}
