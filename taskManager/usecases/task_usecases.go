package usecases

import (
	"taskManager/domain"
	"taskManager/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCases interface {
	CreateTask(input domain.TaskInput) (*domain.Task, error)
	GetTask(id primitive.ObjectID) (*domain.Task, error)
	GetAllTasks() ([]domain.Task, error)
	UpdateTask(input domain.TaskInput, id primitive.ObjectID) (*domain.Task, error)
	DeleteTask(id primitive.ObjectID) error
}

type taskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUseCase(taskRepository repositories.TaskRepository) TaskUseCases {
	return &taskUseCase{taskRepository: taskRepository}
}

func (taskUseCase *taskUseCase) CreateTask(input domain.TaskInput) (*domain.Task, error) {
	task := &domain.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	if err := taskUseCase.taskRepository.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (taskUseCase *taskUseCase) GetTask(id primitive.ObjectID) (*domain.Task, error) {
	return taskUseCase.taskRepository.FindById(id)
}

func (taskUseCase *taskUseCase) GetAllTasks() ([]domain.Task, error) {
	return taskUseCase.taskRepository.FindAll()
}

func (taskUseCase *taskUseCase) UpdateTask(input domain.TaskInput, id primitive.ObjectID) (*domain.Task, error) {
	task := &domain.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	err := taskUseCase.taskRepository.UpdateById(id, task)
	return task, err
}

func (taskUseCase *taskUseCase) DeleteTask(id primitive.ObjectID) error {
	return taskUseCase.taskRepository.Delete(id)
}
