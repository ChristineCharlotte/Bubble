package main

import (
	"Bubble/dao"
	"Bubble/models"
	"Bubble/routers"
)

func main() {
	// 创建数据库 sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":9000")
}
