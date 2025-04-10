package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Connect(databaseName string) error
	Disconnect() error
	Clear() error
	UserOwnsEntity(ctx context.Context, UserID string, EntityID ObjectID, Collection string) error // TODO: Move this to utils

	GetFlowers(ctx context.Context) ([]Flower, error)
	GetUserFlowers(ctx context.Context, userID string) ([]Flower, error)
	GetAllFlowersRelatedToSite(ctx context.Context, siteID ObjectID, userID string) ([]Flower, error)
	AddFlower(ctx context.Context, newFlower Flower) (*Flower, error)
	DeleteFlower(ctx context.Context, id ObjectID) (bool, error)
	ToggleFlowerVisibility(ctx context.Context, userID string, flowerID ObjectID) (*bool, error)
	ModifyFlower(ctx context.Context, id ObjectID, newFlower Flower) (*Flower, error)
	DeleteMultipleFlowers(ctx context.Context, flowerIDs []ObjectID) error
	UpdateVisibilityByTime(ctx context.Context, timestamp time.Time) (modified int64, err error)

	AddSite(ctx context.Context, newSite Site) (*Site, error)
	GetRootSites(ctx context.Context, userID string) ([]Site, error)
	GetSite(ctx context.Context, siteID ObjectID, userID string) (bson.M, error)
	DeleteSite(ctx context.Context, siteID ObjectID, userID string) (*mongoDriver.DeleteResult, error)
	AddFlowerToSite(ctx context.Context, siteID ObjectID, flowerID ObjectID) error
	GetSiteByID(ctx context.Context, siteID ObjectID) (*Site, error)

	AddImage(ctx context.Context, newImage Image) (*Image, error)
	DeleteImage(ctx context.Context, id ObjectID) (bool, error)
	GetImagesByEntity(ctx context.Context, entityID string) ([]Image, error)
	SetFavoriteImage(ctx context.Context, UserID string, EntityID, ImageID ObjectID, Collection string) error
	GetImageByID(ctx context.Context, imageID ObjectID) (*Image, error)
	ClearFavoriteImage(ctx context.Context, UserID string, EntityID ObjectID, Collection string) error
}

type MongoDatabase struct {
	databaseURI string
	client      *mongoDriver.Client
}

type ObjectID = primitive.ObjectID

var NilObjectID ObjectID

var mongoDb *mongoDriver.Database

func NewMongoDatabase(databaseURI string) *MongoDatabase {
	return &MongoDatabase{databaseURI, nil}
}

func (mDb *MongoDatabase) Connect(databaseName string) error {
	timeout := 10 * time.Second
	clientOptions := options.Client().ApplyURI(mDb.databaseURI).SetTimeout(timeout)
	var err error
	mDb.client, err = mongoDriver.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	if err := mDb.client.Ping(context.Background(), nil); err != nil {
		return err
	}

	log.Println("Connected to MongoDB")

	mongoDb = mDb.client.Database(databaseName)

	return nil
}

func (mDb *MongoDatabase) Disconnect() error {
	return mDb.client.Disconnect(context.Background())
}

func (mDb *MongoDatabase) Clear() error {
	return mongoDb.Drop(context.Background())
}

func (mDb MongoDatabase) UserOwnsEntity(ctx context.Context, UserID string, EntityID ObjectID, Collection string) error {
	var user string
	if Collection == "flowers" {
		user = "grower"
	} else {
		user = "owner"
	}
	opts := options.Count().SetLimit(1)
	count, err := mongoDb.Collection(Collection).CountDocuments(
		ctx,
		bson.M{"_id": EntityID, user: UserID},
		opts,
	)
	if err != nil {
		return err
	}

	if count < 1 {
		return fmt.Errorf("user %s does not own %s in %s", UserID, EntityID.Hex(), Collection)
	}

	return nil
}

func ParseID(id string) (ObjectID, error) {
	parsed, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return NilObjectID, fmt.Errorf("error parsing id %q: %w", id, err)
	}

	return parsed, nil
}
