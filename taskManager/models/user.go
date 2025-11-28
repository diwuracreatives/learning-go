package models

import (
	"taskManager/types"

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
