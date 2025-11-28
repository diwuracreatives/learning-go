package main

import (
	"taskManager/data"
	"taskManager/database"
	"taskManager/router"
)

func main() {

	database.Setup()

	taskService := data.NewTaskService()
	userService := data.NewUserService()

	r := router.TaskManagerRouter(taskService, userService)

	err := r.Run("localhost:8000")
	if err != nil {
		return
	}

}
