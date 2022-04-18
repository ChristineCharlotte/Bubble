package controller

import (
	"Bubble/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var err error

var (
	DB *gorm.DB
)

func Create(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	// 2.存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"code": 2000,
	//	"msg": "success",
	//	"data": todo,
	//})

}

// CheckAll 查询todo里的全部数据
func CheckAll(c *gin.Context) {
	todoList, err := models.GetAllTodo()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

}

// Modify 修改
func Modify(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err = models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
