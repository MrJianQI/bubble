package controller

import (
	"bubble/modles"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context)  {
	//前端页面填写代办事项，点击提交，会发请求到这里
	//1. 从请求中把数据拿出来
	var todo modles.Todo
	c.BindJSON(&todo)
	//2. 存入数据库
	//3. 返回响应
	err := modles.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查询todo表里所有的数据
	todoList, err := modles.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo (c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
		return
	}
	todo, err := modles.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = modles.UpdateATodo(todo); err != nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo (c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id不存在"})
		return
	}
	if err :=modles.DeleteATodo(id);err != nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}

}