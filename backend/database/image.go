package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Image struct {
	ID         ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	FileFormat string    `json:"file_format" bson:"file_format"`
	Note       string    `json:"note" form:"note"`
	Entity     *ObjectID `json:"entity" form:"entity"`
	Owner      ObjectID  `json:"owner"`
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

func (mDb MongoDatabase) GetImagesByEntity (ctx context.Context, entityID string) (*[]Image, error) {
	objID, err := ParseID(entityID)
	if err != nil {
		return  nil, err
	}

	err = db.Collection("images").Find(ctx, bson.M{"entity": objID}).Decode(&image)
	if err != nil {
		return nil, err
	}

	images := make([]Image, 0)
	if err := cursor.All(ctx, &images); err != nil {
		return nil, err
	}

	return &images, nil
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
