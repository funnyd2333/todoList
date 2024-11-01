package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v
	v1Group := r.Group("/v1")
	{
		//待办事项

		//添加
		v1Group.POST("/todo", controller.AddTodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.ShowTodos)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
