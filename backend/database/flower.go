package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flower struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	LatinName string             `json:"latin_name" bson:"latin_name"`
	AddedTime time.Time          `json:"added_time" bson:"added_time"`
}

func (aDb ActualDatabase) GetFlowers(ctx context.Context) ([]Flower, error) {
	cursor, err := db.Collection("flowers").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	flowers := make([]Flower, 0)
	if err := cursor.All(ctx, &flowers); err != nil {
		return nil, err
	}

	return flowers, nil
}

func (aDb ActualDatabase) AddFlower(ctx context.Context, newFlower Flower) (*Flower, error) {
	insertResult, err := db.Collection("flowers").InsertOne(ctx, newFlower)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := db.Collection("flowers").FindOne(ctx, filter)

	createdFlower := &Flower{}
	createdRecord.Decode(createdFlower)

	return createdFlower, nil
}

func (aDb ActualDatabase) DeleteFlower(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": objectID}
	result, err := db.Collection("flowers").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, err
}
