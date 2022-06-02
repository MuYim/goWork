package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//连接数据库
func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) //不要加“:”,表示给全局变量赋值，而不是新建
	if err != nil {
		return
	}
	err = DB.DB().Ping() //测试是否连接成功
	return
}

func Close() {
	DB.Close()
}
