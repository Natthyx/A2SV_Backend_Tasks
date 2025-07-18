package router

import (
	"task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	t := r.Group("/tasks")
	{
		t.GET("", controllers.GetTasks)
		t.GET(":id", controllers.GetTask)
		t.POST("", controllers.CreateTask)
		t.PUT(":id", controllers.UpdateTask)
		t.DELETE(":id", controllers.DeleteTask)
	}
	return r
}
