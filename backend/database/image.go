package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Image struct {
	ID     ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Note   string    `json:"note" form:"note"`
	Entity *ObjectID `json:"entity" form:"entity"`
	Owner  ObjectID  `json:"owner"`
}

func (mDb MongoDatabase) AddImage(ctx context.Context, newImage Image) (*Image, error) {
	insertResult, err := db.Collection("images").InsertOne(ctx, newImage)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := db.Collection("images").FindOne(ctx, filter)

	createdImage := &Image{}
	err = createdRecord.Decode(createdImage)
	if err != nil {
		return nil, err
	}

	return createdImage, nil
}

func (mDb MongoDatabase) GetImage(ctx context.Context, imageID ObjectID, userID ObjectID) (*Image, error) {
	resultImage := &Image{}
	filter := bson.M{"_id": imageID, "owner": userID}
	idErr := db.Collection("images").FindOne(ctx, filter).Decode(resultImage)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find image", imageID, "but it doesn't exist")
		}
		return nil, idErr
	}

	return resultImage, nil
}

func (mDb MongoDatabase) DeleteImage(ctx context.Context, id ObjectID) (bool, error) {
	var image Image
	err := db.Collection("images").FindOne(ctx, bson.M{"_id": id}).Decode(&image)
	if err != nil {
		return false, nil
	}

	filter := bson.M{"_id": id}
	result, err := db.Collection("images").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, err
}
