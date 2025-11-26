package data

import (
	"context"
	"errors"
	"log"
	"sync"
	"taskManager/database"
	"taskManager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	m     sync.Mutex
	tasks map[int]models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[int]models.Task),
	}
}

func (taskService *TaskService) CreateTask(input models.TaskInput) models.Task {
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	_, err := database.TaskCollection.InsertOne(context.TODO(), task)

	if err != nil {
		log.Printf("Error inserting task: %v", err)
		return models.Task{}
	}
	return task
}

func (taskService *TaskService) GetTask(id primitive.ObjectID) (models.Task, bool) {
	filter := bson.M{"_id": id}
	result := database.TaskCollection.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return models.Task{}, false
		}
		log.Printf("Error finding task with ID: %v", result.Err())
	}

	var task models.Task
	if err := result.Decode(&task); err != nil {
		log.Printf("Error decoding task: %v", err)
		return models.Task{}, false
	}
	return task, true
}

func (taskService *TaskService) GetAllTasks() []models.Task {
	var tasks []models.Task

	cursor, err := database.TaskCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Printf("Error finding all tasks: %v", err)
		return nil
	}

	defer func() {
		if closeErr := cursor.Close(context.TODO()); closeErr != nil {
			log.Printf("Error closing cursor: %v", closeErr)
		}
	}()

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		log.Printf("Error decoding tasks: %v", err)
		if cursorErr := cursor.Err(); cursorErr != nil {
			log.Printf("Cursor iteration error: %v", cursorErr)
		}
		return nil
	}

	return tasks
}

func (taskService *TaskService) UpdateTask(input models.TaskInput, id primitive.ObjectID) (models.Task, bool) {
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
	}

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": task,
	}

	result, err := database.TaskCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Printf("Error updating task: %v", err)
		return models.Task{}, false
	}

	if result.ModifiedCount == 0 {
		return models.Task{}, false
	}

	return task, true
}

func (taskService *TaskService) DeleteTask(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	deletedTask, err := database.TaskCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return err
	}

	if deletedTask.DeletedCount == 0 {
		return errors.New("task with id not found")
	}

	return nil
}
