package controller

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
url--->logic---->model
请求来了---》控制器----》业务逻辑----》模型层的增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(context *gin.Context) {
	//1、从请求中把数据拿出来
	var todo models.Todo
	context.BindJSON(&todo)
	//2、存入数据库//3、返回响应
	err := models.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(context *gin.Context) {
	//查询表中所有的数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "未填写id"})
		return
	}
	todo, err := models.GetTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	context.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"massage": "delete success"})
	}
}
