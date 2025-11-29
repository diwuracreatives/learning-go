package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func Setup() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, dbErr := mongo.Connect(context.TODO(), clientOptions)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	db = client.Database("taskdb")

	TaskCollection = db.Collection("tasks")
	UserCollection = db.Collection("users")

	if TaskCollection == nil {
		log.Fatal("Failed to create to task collection")
	}

	fmt.Printf("Connected to MongoDB database: %s", TaskCollection.Name())
}

func GetDB() *mongo.Database {
	return db
}
