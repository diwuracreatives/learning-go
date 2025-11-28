package router

import (
	"taskManager/controllers"
	"taskManager/data"
	"taskManager/middleware"

	"github.com/gin-gonic/gin"
)

func TaskManagerRouter(taskService *data.TaskService, userService *data.UserService) *gin.Engine {
	router := gin.Default()
	taskController := controllers.NewTaskController(taskService)
	userController := controllers.NewUserController(userService)

	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)

	router.GET("/tasks/:id", middleware.AuthMiddleware(), taskController.GetTask)
	router.GET("/tasks", middleware.AuthMiddleware(), taskController.GetAllTasks)

	admin := router.Group("/")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/tasks", taskController.CreateTask)
		admin.PUT("/tasks/:id", taskController.UpdateTask)
		admin.DELETE("/tasks/:id", taskController.DeleteTask)
		admin.PATCH("/promote/:id", userController.PromoteUser)

	}

	return router
}
