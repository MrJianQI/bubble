package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//数据库连接
func InitMysql() (err error) {
	dsn := "root:980806@tcp(127.0.0.1:3306)/first?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) //没有:=,是因为是申明一个全局的db变量，不在initMysql里声明，上面的var中做了全局申明
	if err != nil {
		return
	}
	return DB.DB().Ping() //返回数据库的连接
}

func CloseDB() {
	DB.Close()
}
