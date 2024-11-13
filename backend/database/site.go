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
	Name      string      `json:"name"`
	AddedTime time.Time   `json:"added_time" bson:"added_time"`
	Note      string      `json:"note"`
	Parent    *ObjectID   `json:"parent"`
	Flowers   []*ObjectID `json:"flowers"`
	Owner     *ObjectID   `json:"owner"`
}

func (mDb MongoDatabase) AddSite(ctx context.Context, newSite Site) (*Site, error) {
	insertResult, err := db.Collection("sites").InsertOne(ctx, newSite)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := db.Collection("sites").FindOne(ctx, filter)

	createdSite := &Site{}
	err = createdRecord.Decode(createdSite)

	if err != nil {
		return nil, err
	}

	return createdSite, nil
}

func (mDb MongoDatabase) GetRootSites(ctx context.Context, userID ObjectID) ([]Site, error) {
	cursor, err := db.Collection("sites").Find(ctx, bson.M{"parent": nil, "owner": userID})
	if err != nil {
		return nil, err
	}

	var foundSites []Site
	if err := cursor.All(ctx, &foundSites); err != nil {
		return nil, err
	}

	return foundSites, nil
}

// GetSite returns a bson.M containing fields "site", "subSites" and "route", where
// "site" is a Site corresponding to given siteID,
// "subSites" is a list of Sites (_id, name and note) whose parent is "site", and
// "route" is a list of Sites (_id and name) that create a route from "site" (last element is a RootSite).
func (mDb MongoDatabase) GetSite(ctx context.Context, siteID ObjectID, userID ObjectID) (bson.M, error) {
	// Fetch the site
	var resultSite bson.M

	filter := bson.M{"_id": siteID, "owner": userID}
	idErr := db.Collection("sites").FindOne(ctx, filter).Decode(&resultSite)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find site", siteID, "but it doesn't exist")
		}
		return nil, idErr
	}

	// Fetch subsites
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "parent", Value: siteID}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "name", Value: 1}}}}
	unsetStage := bson.D{{Key: "$unset", Value: bson.A{"parent", "addedTime", "owner", "flowers", "added_time"}}}

	cursor, err := db.Collection("sites").Aggregate(ctx, mongo.Pipeline{matchStage, sortStage, unsetStage})
	if err != nil {
		return nil, err
	}

	var subSites []bson.M

	if err = cursor.All(ctx, &subSites); err != nil {
		return nil, err
	}

	// Fetch route to site

	return bson.M{"site": resultSite, "subsites": subSites}, nil
}

func (mDb MongoDatabase) DeleteSite(ctx context.Context, siteID ObjectID, userID ObjectID) (*mongo.DeleteResult, error) {
	// Start pipeline with top level parent Site
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: siteID},
			{Key: "owner", Value: userID},
		}},
	}
	// Search for all children and their children
	graphLookupStage := bson.D{
		{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: "sites"},
			{Key: "startWith", Value: "$_id"},
			{Key: "connectFromField", Value: "_id"},
			{Key: "connectToField", Value: "parent"},
			{Key: "as", Value: "related"},
		}},
	}
	// Strip down everything except _id for each child Site
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "id", Value: "$related._id"},
		}},
	}
	// Open up array of documents to a stream of documents
	unwindStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$id"},
		},
		}}

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
			return nil, errors.New("fetched sub site ID was of wrong type")
		}
		ids = append(ids, sub_id)
	}

	log.Println("DELETE sites", ids)

	flowerDeleteFilter := bson.M{"site": bson.M{"$in": ids}}
	deleteFlowerResult, err := db.Collection("flowers").DeleteMany(ctx, flowerDeleteFilter)
	if err != nil {
		return nil, err
	}
	log.Println("Deleted flowers count:", deleteFlowerResult.DeletedCount)

	deleteFilter := bson.M{"_id": bson.M{"$in": ids}}
	deleteResult, err := db.Collection("sites").DeleteMany(ctx, deleteFilter)
	if err != nil {
		log.Println("DELETE FAILED: ", err)
		return nil, err
	}

	return deleteResult, nil
}

func (mDb MongoDatabase) AddFlowerToSite(ctx context.Context, siteID ObjectID, flowerID ObjectID) error {
	update := bson.M{"$push": bson.M{"flowers": flowerID}}
	_, err := db.Collection("sites").UpdateOne(ctx, bson.M{"_id": siteID}, update)
	return err
}

func (mDb MongoDatabase) GetSiteByID(ctx context.Context, siteID ObjectID) (*Site, error) {
	var site Site
	err := db.Collection("sites").FindOne(ctx, bson.M{"_id": siteID}).Decode(&site)
	return &site, err
}
