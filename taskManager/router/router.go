package router

import (
	"taskManager/controllers"
	"taskManager/data"

	"github.com/gin-gonic/gin"
)

func TaskManagerRouter(taskService *data.TaskService) *gin.Engine {
	router := gin.Default()
	taskController := controllers.NewTaskController(taskService)

	router.GET("/tasks/:id", taskController.GetTask)
	router.GET("/tasks", taskController.GetAllTasks)
	router.POST("/tasks", taskController.CreateTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	return router
}
