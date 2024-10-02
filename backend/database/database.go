package database

import (
	"context"
	"log"
	"time"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	CountUsersWithEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, newUser User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)

	GetFlowers(ctx context.Context) ([]Flower, error)
    AddFlower(ctx context.Context, newFlower Flower) (*Flower, error)
    DeleteFlower(ctx context.Context, id string) (bool, error)

	AddSite(ctx context.Context, newSite Site) (*Site, error)
	GetRootSites(ctx context.Context) ([]Site, error)
	GetSite(ctx context.Context, id string) (bson.M, error)
	DeleteSite(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type MongoDatabase struct {}

type MockDatabase struct {
	mock.Mock
}

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
