package repositories

import (
	"context"
	"errors"
	"log"
	"taskManager/database"
	"taskManager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	FindById(id primitive.ObjectID) (*domain.Task, error)
	FindAll() ([]domain.Task, error)
	UpdateById(id primitive.ObjectID, task *domain.Task) error
	Delete(id primitive.ObjectID) error
}

type taskRepo struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
	return &taskRepo{collection: db.Collection("tasks")}
}

func (r *taskRepo) Create(task *domain.Task) error {
	_, insertTaskErr := database.TaskCollection.InsertOne(context.TODO(), task)
	return insertTaskErr
}

func (r *taskRepo) FindById(id primitive.ObjectID) (*domain.Task, error) {
	filter := bson.M{"_id": id}
	result := database.TaskCollection.FindOne(context.TODO(), filter)

	var task domain.Task
	if err := result.Decode(&task); err != nil {
		log.Printf("Error decoding task: %v", err)
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) FindAll() ([]domain.Task, error) {
	cursor, err := database.TaskCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Printf("Error finding all tasks: %v", err)
		return nil, err
	}

	defer func() {
		if closeErr := cursor.Close(context.TODO()); closeErr != nil {
			log.Printf("Error closing cursor: %v", closeErr)
		}
	}()

	var tasks []domain.Task

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		log.Printf("Error decoding tasks: %v", err)
		if cursorErr := cursor.Err(); cursorErr != nil {
			log.Printf("Cursor iteration error: %v", cursorErr)
		}
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepo) UpdateById(id primitive.ObjectID, task *domain.Task) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": task,
	}

	result, err := database.TaskCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Printf("Error updating task: %v", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepo) Delete(id primitive.ObjectID) error {
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
