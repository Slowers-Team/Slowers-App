package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flower struct {
	ID          ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name"`
	LatinName   string    `json:"latin_name" bson:"latin_name"`
	AddedTime   time.Time `json:"added_time" bson:"added_time"`
	Grower      *ObjectID `json:"grower"`
	GrowerEmail string    `json:"grower_email" bson:"grower_email"`
	Site        *ObjectID `json:"site"`
	SiteName    string    `json:"site_name" bson:"site_name"`
}

func (mDb MongoDatabase) GetFlowers(ctx context.Context) ([]Flower, error) {
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

func (mDb MongoDatabase) GetUserFlowers(ctx context.Context, userID ObjectID) ([]Flower, error) {
	cursor, err := db.Collection("flowers").Find(ctx, bson.M{"grower": userID})
	if err != nil {
		return nil, err
	}

	flowers := make([]Flower, 0)
	if err := cursor.All(ctx, &flowers); err != nil {
		return nil, err
	}

	return flowers, nil
}

func (mDb MongoDatabase) AddFlower(ctx context.Context, newFlower Flower) (*Flower, error) {
	insertResult, err := db.Collection("flowers").InsertOne(ctx, newFlower)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := db.Collection("flowers").FindOne(ctx, filter)

	createdFlower := &Flower{}
	err = createdRecord.Decode(createdFlower)
	if err != nil {
		return nil, err
	}

	return createdFlower, nil
}

func (mDb MongoDatabase) DeleteFlower(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	var flower Flower
	err = db.Collection("flowers").FindOne(ctx, bson.M{"_id": objectID}).Decode(&flower)
	if err != nil {
		return false, nil
	}

	if flower.Site != nil {
		update := bson.M{"$pull": bson.M{"flowers": objectID}}
		_, err = db.Collection("sites").UpdateOne(ctx, bson.M{"_id": flower.Site}, update)
		if err != nil {
			return true, err
		}
	}

	filter := bson.M{"_id": objectID}
	result, err := db.Collection("flowers").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, err
}
