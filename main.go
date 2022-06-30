package main

import (
	"bubble/dao"
	"bubble/routers"
)

func main() {
	//创建数据库
	// sql: CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.CloseDB() // 程序退出关闭数据库连接
	//模型绑定,也是建表
	dao.InitMysql()
	//注册路由
	r := routers.SetupRputer()
	r.Run()
	//err = modles.DeleteATodo("1")

}
