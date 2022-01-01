package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	logging "github.com/killtheverse/nitd-results/app/logging"
)

// SetIndexes will set the indexes in database
func SetIndexes(collection *mongo.Collection, keys bsonx.Doc) {
	index := mongo.IndexModel{}
	index.Keys = keys
	unique := true
	index.Options = &options.IndexOptions{
		Unique: &unique,
	}
	opts := options.CreateIndexes().SetMaxTime(10*time.Second)
	_, err := collection.Indexes().CreateOne(context.Background(), index, opts)
	if err != nil {
		logging.Fatal("[ERROR]: While creating indexes. %v", err)
	}
}