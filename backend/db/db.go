package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseClient = mongo.Client
type ObjectID primitive.ObjectID

var db *mongo.Database

func Connect(mongoURI string, databaseName string) (*DatabaseClient, error) {
	timeout := 10 * time.Second
	clientOptions := options.Client().ApplyURI(mongoURI).SetTimeout(timeout)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")

	db = client.Database(databaseName)

	return client, nil
}

func Disconnect(client *DatabaseClient) error {
	return client.Disconnect(context.Background())
}

func Clear() error {
	return db.Drop(context.Background())
}

func IsValidID(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}
