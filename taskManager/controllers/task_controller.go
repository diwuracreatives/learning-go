package controllers

import (
	"net/http"
	"strings"
	"taskManager/data"
	"taskManager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskService *data.TaskService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (taskController *TaskController) CreateTask(c *gin.Context) {
	var newTask models.TaskInput

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

	task := taskController.taskService.CreateTask(newTask)
	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func (taskController *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, exists := taskController.taskService.GetTask(id)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with id " + idParam + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (taskController *TaskController) GetAllTasks(c *gin.Context) {
	tasks := taskController.taskService.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (taskController *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var taskInput models.TaskInput

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

	task, exists := taskController.taskService.UpdateTask(taskInput, id)
	if !exists {
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

	err = taskController.taskService.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with id " + idParam + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task with id " + idParam + " deleted"})
}
