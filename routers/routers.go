package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRputer() *gin.Engine {
	r := gin.Default()
	//设置模板文件引用的静态文件位置
	r.Static("/static", "static")
	//设置模板文件位置
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//路由组，所有的路径前加一个v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有代办事项
		v1Group.GET("/todo",controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
