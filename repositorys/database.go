package repositorys

import (
	"KiKo/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

// 全局DB对象
var Db *gorm.DB

func init() {
	// 数据库
	dbType := conf.Viper.GetString("database.driver")
	addr := conf.Viper.GetString("mysql.addr")
	port := conf.Viper.GetString("mysql.port")
	name := conf.Viper.GetString("mysql.name")
	params := conf.Viper.GetString("mysql.params")
	user := conf.Viper.GetString("mysql.user")
	pass := conf.Viper.GetString("mysql.pass")
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", user, pass, addr, port, name, params)

	var err error
	Db, err = gorm.Open(dbType, url)
	if err != nil {
		log.Printf("Open mysql failed,err:%v\n", err)
		panic(err)
	}
	Db.SingularTable(true)
	//启用日志程序，显示详细日志
	Db.LogMode(true)
	Db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	// 连接池设置
	Db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	Db.DB().SetMaxOpenConns(30)                   //最大连接数
	Db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时

}
