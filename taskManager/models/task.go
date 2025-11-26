package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DueDate     time.Time          `bson:"due_date"`
	Status      string             `bson:"status"`
}

type TaskInput struct {
	Title       string `bson:"title" binding:"required"`
	Description string `bson:"description" binding:"required"`
	Status      string `bson:"status" binding:"required"`
}
