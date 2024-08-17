package routes

import (
	"aristio-sagala-test/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("", controllers.CreateTask)
		taskRoutes.GET("", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTask)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}

	return r
}
