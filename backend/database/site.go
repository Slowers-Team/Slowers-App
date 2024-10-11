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

func (mDb MongoDatabase) GetSite(ctx context.Context, id string, userID ObjectID) (bson.M, error) {
	siteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var resultSite bson.M

	filter := bson.M{"_id": siteID, "owner": userID}
	idErr := db.Collection("sites").FindOne(ctx, filter).Decode(&resultSite)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find site", siteID, "but it doesn't exist")
		}
		return nil, idErr
	}

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

	return bson.M{"site": resultSite, "subsites": subSites}, nil
}

func (mDb MongoDatabase) DeleteSite(ctx context.Context, id string, userID ObjectID) (*mongo.DeleteResult, error) {
	siteID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

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
			return nil, errors.New("Fetched sub site ID was of wrong type")
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

func (mDb MongoDatabase) GetAllFlowersRelatedToSite(ctx context.Context, siteID string, userID ObjectID) ([]Flower, error) {
	parentSiteID, err := primitive.ObjectIDFromHex(siteID)
	if err != nil {
		return nil, err
	}

	// Start pipeline with top level parent Site
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: parentSiteID},
			{Key: "owner", Value: userID},
		}},
	}
	// Search for all subsites and their subsites
	graphLookupStage := bson.D{
		{Key: "$graphLookup", Value: bson.D{
			{Key: "from", Value: "sites"},
			{Key: "startWith", Value: "$_id"},
			{Key: "connectFromField", Value: "_id"},
			{Key: "connectToField", Value: "parent"},
			{Key: "as", Value: "related"},
		}},
	}
	// Strip down everything except _id for each subsite
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "id", Value: "$related._id"},
		}},
	}
	// Add parent site to the ids
	concatStage := bson.D{
		{"$addFields", bson.D{
			{"id", bson.D{
				{"$concatArrays", bson.A{bson.A{parentSiteID}, "$id"}},
			}},
		}},
	}
	// Open up array of Sites to a stream of Sites
	unwindSitesStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$id"},
		},
		}}

	// connect each site ID to a list of flowers
	lookupStage := bson.D{
		{"$lookup", bson.D{
			{"from", "flowers"},
			{"localField", "id"},
			{"foreignField", "site"},
			{"as", "flowers"},
		}},
	}
	// Open up arrays of Flowers to a single stream of Flowers
	unwindFlowersStage := bson.D{
		{"$unwind", bson.D{
			{"path", "$flowers"},
		}},
	}
	// Raise each Flower document to root level (instead of being behind "flowers" field)
	replaceRootStage := bson.D{
		{"$replaceRoot", bson.D{
			{"newRoot", "$flowers"},
		}},
	}

	cursor, err := db.Collection("sites").Aggregate(ctx, mongo.Pipeline{
		matchStage, graphLookupStage, projectStage, concatStage, unwindSitesStage,
		lookupStage, unwindFlowersStage, replaceRootStage})
	if err != nil {
		return nil, err
	}

	var flowers []Flower
	if err := cursor.All(ctx, &flowers); err != nil {
		return nil, err
	}

	return flowers, nil
}
