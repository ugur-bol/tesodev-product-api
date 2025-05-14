package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	DB = client
	log.Println("Connected to MongoDB")
}
