package domain

import (
	"taskManager/types"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     types.Role         `bson:"role"`
}

type UserInput struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

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
