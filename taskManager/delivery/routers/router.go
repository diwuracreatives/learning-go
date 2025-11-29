package routers

import (
	"taskManager/delivery/controllers"
	"taskManager/infrastructure"
	"taskManager/usecases"

	"github.com/gin-gonic/gin"
)

func TaskManagerRouter(taskUseCases usecases.TaskUseCases, userUseCases usecases.UserUseCases) *gin.Engine {
	router := gin.Default()
	taskController := controllers.NewTaskController(taskUseCases)
	userController := controllers.NewUserController(userUseCases)

	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)

	router.GET("/tasks/:id", infrastructure.AuthMiddleware(), taskController.GetTask)
	router.GET("/tasks", infrastructure.AuthMiddleware(), taskController.GetAllTasks)

	admin := router.Group("/")
	admin.Use(infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware())
	{
		admin.POST("/tasks", taskController.CreateTask)
		admin.PUT("/tasks/:id", taskController.UpdateTask)
		admin.DELETE("/tasks/:id", taskController.DeleteTask)
		admin.PATCH("/promote/:id", userController.PromoteUser)

	}

	return router
}
