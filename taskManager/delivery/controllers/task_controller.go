package controllers

import (
	"net/http"
	"strings"
	"taskManager/domain"
	"taskManager/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskUseCases usecases.TaskUseCases
}

func NewTaskController(taskUseCases usecases.TaskUseCases) *TaskController {
	return &TaskController{taskUseCases: taskUseCases}
}

func (taskController *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.TaskInput

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request payload invalid",
			"message": err.Error(),
		})
		return
	}

	newTask.Title = strings.TrimSpace(newTask.Title)
	newTask.Description = strings.TrimSpace(newTask.Description)
	newTask.Status = strings.TrimSpace(newTask.Status)

	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	} else if newTask.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
		return
	} else if newTask.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	task, err := taskController.taskUseCases.CreateTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func (taskController *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := taskController.taskUseCases.GetTask(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with id " + idParam + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (taskController *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := taskController.taskUseCases.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (taskController *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var taskInput domain.TaskInput

	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request payload invalid",
			"message": err.Error(),
		})
		return
	}

	taskInput.Title = strings.TrimSpace(taskInput.Title)
	taskInput.Description = strings.TrimSpace(taskInput.Description)
	taskInput.Status = strings.TrimSpace(taskInput.Status)

	if taskInput.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	} else if taskInput.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "description is required"})
		return
	} else if taskInput.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}

	task, err := taskController.taskUseCases.UpdateTask(taskInput, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with id " + idParam + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (taskController *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = taskController.taskUseCases.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with id " + idParam + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task with id " + idParam + " deleted"})
}
