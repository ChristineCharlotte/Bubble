package routers

import (
	"Bubble/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉 gin 框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.Create)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.CheckAll)

		// 查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
		})

		// 修改某一个代办事项
		v1Group.PUT("/todo/:id", controller.Modify)

		// 删除某一个代办事项
		v1Group.DELETE("/todo/:id", controller.Delete)
	}

	return r

}
