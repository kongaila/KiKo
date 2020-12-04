package conf

import (
	"github.com/spf13/viper"
)

// 读取的路径, 默认dev
// 生产 prod
// 测试 test
var path = "dev"

// viper 全局实例
var Viper *viper.Viper

// 初始化viper
func init() {
	Viper = viper.New()
	// 指定配置文件的文件名称
	Viper.SetConfigName("config")
	Viper.AddConfigPath("./conf/" + path)
	// 设置默认的数据源类型
	Viper.SetDefault("database.driver", "mysql")
	err := Viper.ReadInConfig()
	if err != nil {
		// 读取配置文件失败致命错误
		panic(err)
	}
}
