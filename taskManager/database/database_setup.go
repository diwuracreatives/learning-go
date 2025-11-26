package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func Setup() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, dbErr := mongo.Connect(context.TODO(), clientOptions)

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	TaskCollection = client.Database("taskdb").Collection("tasks")

	if TaskCollection == nil {
		log.Fatal("Failed to create to task collection")
	}

	fmt.Printf("Connected to MongoDB database: %s", TaskCollection.Name())
}
