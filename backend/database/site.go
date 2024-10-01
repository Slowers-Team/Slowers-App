package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Site struct {
	ID        ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                `json:"name"`
	AddedTime time.Time             `json:"added_time" bson:"added_time"`
	Note      string                `json:"note"`
	Parent    *ObjectID   `json:"parent"`
	Flowers   []*ObjectID `json:"flowers"`
	Owner     *ObjectID   `json:"owner"`
}

func (aDb ActualDatabase) AddSite(ctx context.Context, newSite Site) (*Site, error) {
	insertResult, err := db.Collection("sites").InsertOne(ctx, newSite)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := db.Collection("sites").FindOne(ctx, filter)

	createdSite := &Site{}
	createdRecord.Decode(createdSite)

	return createdSite, nil
}

func (aDb ActualDatabase) GetRootSites(ctx context.Context) ([]Site, error) {
	cursor, err := db.Collection("sites").Find(ctx, bson.M{"parent": nil})
	if err != nil {
		return nil, err
	}

	var foundSites []Site
	if err := cursor.All(ctx, &foundSites); err != nil {
		return nil, err
	}

	return foundSites, nil
}

func (aDb ActualDatabase) GetSite(ctx context.Context, id string) (bson.M, error) {
	siteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var resultSite bson.M

	filter := bson.M{"_id": siteID}
	idErr := db.Collection("sites").FindOne(ctx, filter).Decode(&resultSite)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find site", siteID, "but it doesn't exist")
		}
		return nil, idErr
	}

	log.Println("found site:", resultSite)

	matchStage := bson.D{{"$match", bson.D{{"parent", siteID}}}}
	sortStage := bson.D{{"$sort", bson.D{{"name", 1}}}}
	unsetStage := bson.D{{"$unset", bson.A{"parent", "addedTime", "owner", "flowers", "added_time"}}}

	cursor, err := db.Collection("sites").Aggregate(ctx, mongo.Pipeline{matchStage, sortStage, unsetStage})
	if err != nil {
		return nil, err
	}

	var subSites []bson.M

	if err = cursor.All(ctx, &subSites); err != nil {
		return nil, err
	}

	log.Println("subsites:", subSites)

	return bson.M{"site": resultSite, "subsites": subSites}, nil
}

func (aDb ActualDatabase) DeleteSite(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	siteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Start pipeline with top level parent Site
	matchStage := bson.D{
		{"$match", bson.D{
			{"_id", siteID},
		}},
	}
	// Search for all children and their children
	graphLookupStage := bson.D{
		{"$graphLookup", bson.D{
			{"from", "sites"},
			{"startWith", "$_id"},
			{"connectFromField", "_id"},
			{"connectToField", "parent"},
			{"as", "related"},
		}},
	}
	// Open up array of documents to a stream of documents
	unwindStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$id"},
		},
		}}
	// Strip down everything except _id for each child Site
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"id", "$related._id"},
		}},
	}

	cursor, err := db.Collection("sites").Aggregate(ctx, mongo.Pipeline{matchStage, graphLookupStage, projectStage, unwindStage})
	if err != nil {
		return nil, err
	}

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	ids := make([]primitive.ObjectID, 0)
	ids = append(ids, siteID)

	for _, res := range result {
		sub_id, err := res["id"].(primitive.ObjectID)
		if !err {
			return nil, errors.New("Fetched sub site ID was of wrong type")
		}
		ids = append(ids, sub_id)
	}

	log.Println("DELETE sites", ids)

	deleteFilter := bson.M{"_id": bson.M{"$in": ids}}
	deleteResult, err := db.Collection("sites").DeleteMany(ctx, deleteFilter)
	if err != nil {
		log.Println("DELETE FAILED: ", err)
		return nil, err
	}

	return deleteResult, nil
}
