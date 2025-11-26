package main

import (
	"taskManager/data"
	"taskManager/database"
	"taskManager/router"
)

func main() {

	database.Setup()

	taskService := data.NewTaskService()

	r := router.TaskManagerRouter(taskService)

	err := r.Run("localhost:8000")
	if err != nil {
		return
	}

}
