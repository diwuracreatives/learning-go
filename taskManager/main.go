package main

import (
	"taskManager/data"
	"taskManager/router"
)

func main() {
	taskService := data.NewTaskService()

	r := router.TaskManagerRouter(taskService)

	err := r.Run("localhost:8000")
	if err != nil {
		return
	}
}
