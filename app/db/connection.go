package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	logger "github.com/killtheverse/nitd-results/app/logging"
)

// Connect will connect to MongoDB Database and return the client
func Connect(mongoURI string) *mongo.Client {
	logger.Write("Connecting to database")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Fatal("[ERROR]: Can't connect to Database: %v\n", err)
	} else {
		logger.Write("Connected to database")
	}
	return client
}

// Disconnect will disconnect the client from the Database
func Disconnect(client *mongo.Client) {
	logger.Write("Disconnecting from database")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		logger.Write("[ERROR]: Error in disconnection: %v\n", err)
	} else {
		logger.Write("Disconnected from database")
	}
}