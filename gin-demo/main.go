package main

import (
	"gin-demo/dao"
	"gin-demo/routers"
)

func main() {
	//创建数据库
	//create database bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出时关闭数据库连接

	//绑定路由
	r := routers.SetupRouter()
	r.Run(":8080")
}
