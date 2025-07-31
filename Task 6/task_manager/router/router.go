package router

import (
	"task_manager/controllers"
	"task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	tasks := r.Group("/tasks")
	tasks.Use(middleware.JWTAuthMiddleware())
	{
		tasks.GET("", controllers.GetTasks)
		tasks.GET(":id", controllers.GetTask)
		tasks.POST("", controllers.CreateTask)
		tasks.PUT(":id", controllers.UpdateTask)
		tasks.DELETE(":id", controllers.DeleteTask)
	}

	return r
}
