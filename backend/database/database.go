package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Connect(databaseName string) error
	Disconnect() error
	Clear() error

	CountUsersWithEmail(ctx context.Context, email string) (int64, error)
	CreateUser(ctx context.Context, newUser User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userID ObjectID) (*User, error)
	SetUserRole(ctx context.Context, userID ObjectID, role string) error

	GetFlowers(ctx context.Context) ([]Flower, error)
	GetUserFlowers(ctx context.Context, userID ObjectID) ([]Flower, error)
	GetAllFlowersRelatedToSite(ctx context.Context, siteID ObjectID, userID ObjectID) ([]Flower, error)
	AddFlower(ctx context.Context, newFlower Flower) (*Flower, error)
	DeleteFlower(ctx context.Context, id ObjectID) (bool, error)
	GetFlowerByID(ctx context.Context, flowerID ObjectID) (*Flower, error)

	AddSite(ctx context.Context, newSite Site) (*Site, error)
	GetRootSites(ctx context.Context, userID ObjectID) ([]Site, error)
	GetSite(ctx context.Context, siteID ObjectID, userID ObjectID) (bson.M, error)
	DeleteSite(ctx context.Context, siteID ObjectID, userID ObjectID) (*mongo.DeleteResult, error)
	AddFlowerToSite(ctx context.Context, siteID ObjectID, flowerID ObjectID) error
	GetSiteByID(ctx context.Context, siteID ObjectID) (*Site, error)
}

type MongoDatabase struct {
	databaseURI string
	client      *mongo.Client
}

type ObjectID = primitive.ObjectID

var NilObjectID ObjectID

var db *mongo.Database

func NewMongoDatabase(databaseURI string) *MongoDatabase {
	return &MongoDatabase{databaseURI, nil}
}

func (mDb *MongoDatabase) Connect(databaseName string) error {
	timeout := 10 * time.Second
	clientOptions := options.Client().ApplyURI(mDb.databaseURI).SetTimeout(timeout)
	var err error
	mDb.client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	if err := mDb.client.Ping(context.Background(), nil); err != nil {
		return err
	}

	log.Println("Connected to MongoDB")

	db = mDb.client.Database(databaseName)

	return nil
}

func (mDb *MongoDatabase) Disconnect() error {
	return mDb.client.Disconnect(context.Background())
}

func (mDb *MongoDatabase) Clear() error {
	return db.Drop(context.Background())
}

func ParseID(id string) (ObjectID, error) {
	parsed, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return NilObjectID, fmt.Errorf("error parsing id %q: %w", id, err)
	}

	return parsed, nil
}
