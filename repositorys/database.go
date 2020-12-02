package repositorys

import (
	"QiqiLike/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// 全局DB对象
var Db *gorm.DB

func init() {
	dbType := conf.Viper.GetString("database.driver")
	host := conf.Viper.GetString("mysql.host")
	port := conf.Viper.GetString("mysql.port")
	name := conf.Viper.GetString("mysql.name")
	params := conf.Viper.GetString("mysql.params")
	user := conf.Viper.GetString("mysql.user")
	pass := conf.Viper.GetString("mysql.pass")
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", user, pass, host, port, name, params)

	var err error
	Db, err = gorm.Open(dbType, url)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
	Db.SingularTable(true)
	//启用日志程序，显示详细日志
	Db.LogMode(true)

	Db.LogMode(true)
	Db.SetLogger(log.New(os.Stdout, "\r\n", 0))

}
