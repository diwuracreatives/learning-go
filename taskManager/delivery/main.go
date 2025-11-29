package main

import (
	"taskManager/database"
	"taskManager/delivery/routers"
	"taskManager/infrastructure"
	"taskManager/repositories"
	"taskManager/usecases"
)

func main() {

	database.Setup()

	userRepo := repositories.NewUserRepository(database.GetDB())
	taskRepo := repositories.NewTaskRepository(database.GetDB())

	taskService := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(infrastructure.NewJwtService(), userRepo, infrastructure.NewPasswordService())

	r := routers.TaskManagerRouter(taskService, userUseCase)

	err := r.Run("localhost:8000")
	if err != nil {
		return
	}

}
