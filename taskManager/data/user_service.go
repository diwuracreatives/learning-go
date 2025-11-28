package data

import (
	"context"
	"errors"
	"log"
	"sync"
	"taskManager/database"
	"taskManager/models"
	"taskManager/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	m sync.Mutex
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) GetUserByUsername(username string) (models.User, error) {
	filter := bson.M{"username": username}
	result := database.UserCollection.FindOne(context.TODO(), filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			log.Printf("an error occured: %v", result.Err())
			return models.User{}, errors.New("user with username not found")
		}
		log.Printf("an error occured: %v", result.Err())
	}

	var user models.User
	if err := result.Decode(&user); err != nil {
		log.Printf("Error decoding user: %v", err)
		return models.User{}, err
	}
	return user, nil

}

func (userService *UserService) CreateUser(input models.UserInput) (models.User, error) {
	existingUser, _ := userService.GetUserByUsername(input.Username)

	if existingUser.Username != "" {
		return models.User{}, errors.New("user with username already exists")
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Role:     types.User,
	}

	count, err := database.UserCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("an error occured getting user collection count: %v", err)
	}

	if count == 0 {
		user.Role = types.Admin
	}

	// hash the password
	hashedPassword, passwordErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if passwordErr != nil {
		log.Printf("an error occured generating password: %v", err)
		return models.User{}, errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	_, insertUserErr := database.UserCollection.InsertOne(context.TODO(), user)

	if insertUserErr != nil {
		log.Printf("Error inserting user: %v", err)
		return models.User{}, errors.New("internal server error")
	}
	return user, nil
}

func (userService *UserService) LoginUser(input models.UserInput) (models.User, error) {
	user, err := userService.GetUserByUsername(input.Username)
	if err != nil || user.Username == "" {
		return models.User{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}

func (userService *UserService) PromoteUser(id primitive.ObjectID) error {

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
