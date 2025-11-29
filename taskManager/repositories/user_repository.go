package repositories

import (
	"context"
	"errors"
	"log"
	"taskManager/database"
	"taskManager/domain"
	"taskManager/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindUserByUsername(username string) (*domain.User, error)
	Count() (int64, error)
	UpdateById(id primitive.ObjectID) error
}

type userRepo struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepo{collection: db.Collection("users")}
}

func (r *userRepo) Create(user *domain.User) error {
	_, insertUserErr := database.UserCollection.InsertOne(context.TODO(), user)
	return insertUserErr
}

func (r *userRepo) FindUserByUsername(username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	result := database.UserCollection.FindOne(context.TODO(), filter)

	var user domain.User
	if err := result.Decode(&user); err != nil {
		log.Printf("Error decoding user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Count() (int64, error) {
	count, err := database.UserCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("an error occured getting user collection count: %v", err)
	}
	return count, nil
}

func (r *userRepo) UpdateById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{"role": types.Admin},
	}

	result, err := database.UserCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Printf("Error updating user role: %v", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("user with username not found")
	}

	return nil
}
