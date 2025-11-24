package data

import (
	"errors"
	"sync"
	"taskManager/models"
)

type TaskService struct {
	m      sync.Mutex
	taskId int
	tasks  map[int]models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:  make(map[int]models.Task),
		taskId: 1,
	}
}

func (taskService *TaskService) CreateTask(input models.TaskInput) models.Task {
	taskService.m.Lock()
	defer taskService.m.Unlock()

	task := models.Task{
		ID:          taskService.taskId,
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	taskService.tasks[task.ID] = task
	taskService.taskId++
	return task
}

func (taskService *TaskService) GetTask(id int) (models.Task, bool) {
	taskService.m.Lock()
	defer taskService.m.Unlock()

	task, ok := taskService.tasks[id]
	if !ok {
		return models.Task{}, false
	}
	return task, ok
}

func (taskService *TaskService) GetAllTasks() []models.Task {
	tasks := make([]models.Task, 0)
	taskService.m.Lock()
	defer taskService.m.Unlock()

	for _, task := range taskService.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (taskService *TaskService) UpdateTask(input models.TaskInput, id int) (models.Task, bool) {
	taskService.m.Lock()
	defer taskService.m.Unlock()

	task, ok := taskService.tasks[id]
	if !ok {
		return models.Task{}, false
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status

	taskService.tasks[id] = task

	return task, true
}

func (taskService *TaskService) DeleteTask(id int) error {
	taskService.m.Lock()
	defer taskService.m.Unlock()

	_, ok := taskService.tasks[id]
	if !ok {
		return errors.New("task with id not found")
	}

	delete(taskService.tasks, id)
	return nil
}
