package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Flower struct {
	ID            ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string    `json:"name"`
	LatinName     string    `json:"latin_name" bson:"latin_name"`
	AddedTime     time.Time `json:"added_time" bson:"added_time"`
	Grower        *string   `json:"grower"`
	GrowerEmail   string    `json:"grower_email" bson:"grower_email"`
	Site          *ObjectID `json:"site"`
	SiteName      string    `json:"site_name" bson:"site_name"`
	Quantity      int       `json:"quantity"`
	Visible       bool      `json:"visible" bson:"visible"`
	FavoriteImage string    `json:"favorite_image" bson:"favorite_image"`
}

func (mDb MongoDatabase) GetFlowers(ctx context.Context) ([]Flower, error) {
	cursor, err := mongoDb.Collection("flowers").Find(ctx, bson.M{"visible": true})

	if err != nil {
		return nil, err
	}

	flowers := make([]Flower, 0)
	if err := cursor.All(ctx, &flowers); err != nil {
		return nil, err
	}

	return flowers, nil
}

func (mDb MongoDatabase) GetUserFlowers(ctx context.Context, userID string) ([]Flower, error) {
	cursor, err := mongoDb.Collection("flowers").Find(ctx, bson.M{"grower": userID})
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
	insertResult, err := mongoDb.Collection("flowers").InsertOne(ctx, newFlower)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := mongoDb.Collection("flowers").FindOne(ctx, filter)

	createdFlower := &Flower{}
	err = createdRecord.Decode(createdFlower)
	if err != nil {
		return nil, err
	}

	return createdFlower, nil
}

func (mDb MongoDatabase) DeleteFlower(ctx context.Context, id ObjectID) (bool, error) {
	var flower Flower
	err := mongoDb.Collection("flowers").FindOne(ctx, bson.M{"_id": id}).Decode(&flower)
	if err != nil {
		return false, nil
	}

	if flower.Site != nil {
		update := bson.M{"$pull": bson.M{"flowers": id}}
		_, err = mongoDb.Collection("sites").UpdateOne(ctx, bson.M{"_id": flower.Site}, update)
		if err != nil {
			return true, err
		}
	}

	filter := bson.M{"_id": id}
	result, err := mongoDb.Collection("flowers").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, err
}

func (mDb MongoDatabase) GetAllFlowersRelatedToSite(ctx context.Context, siteID ObjectID, userID string) ([]Flower, error) {
	// Start pipeline with top level parent Site
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "_id", Value: siteID},
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
		{Key: "$addFields", Value: bson.D{
			{Key: "id", Value: bson.D{
				{Key: "$concatArrays", Value: bson.A{bson.A{siteID}, "$id"}},
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
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "flowers"},
			{Key: "localField", Value: "id"},
			{Key: "foreignField", Value: "site"},
			{Key: "as", Value: "flowers"},
		}},
	}
	// Open up arrays of Flowers to a single stream of Flowers
	unwindFlowersStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$flowers"},
		}},
	}
	// Raise each Flower document to root level (instead of being behind "flowers" field)
	replaceRootStage := bson.D{
		{Key: "$replaceRoot", Value: bson.D{
			{Key: "newRoot", Value: "$flowers"},
		}},
	}

	cursor, err := mongoDb.Collection("sites").Aggregate(ctx, mongo.Pipeline{
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

// ToggleFlowerVisibility sets the toggles (false->true or true->false) flower's visibility,
// and returns the new value or an error.
// Visibility can be set if flower has at least one image attached.
func (mDb MongoDatabase) ToggleFlowerVisibility(ctx context.Context, userID string, flowerID ObjectID) (*bool, error) {
	opts := options.Count().SetLimit(1)
	count, err := mongoDb.Collection("images").CountDocuments(
		ctx,
		bson.M{"entity": flowerID},
		opts,
	)
	if err != nil {
		return nil, err
	}

	if count < 1 {
		return nil, fmt.Errorf("no image attached to flower %s", flowerID.Hex())
	}

	filter := bson.M{"_id": flowerID}
	update := bson.A{bson.M{"$set": bson.M{"visible": bson.M{"$not": "$visible"}}}}
	updateOpts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.M{"_id": 0, "visible": 1})

	var updatedVisibility bson.M
	err = mongoDb.Collection("flowers").FindOneAndUpdate(ctx, filter, update, updateOpts).Decode(&updatedVisibility)

	if err != nil {
		return nil, err
	}

	ret := updatedVisibility["visible"].(bool)

	if ret {
		newTime := time.Now()
		updateTime := bson.A{bson.M{"$set": bson.M{"added_time": newTime}}}
		var updatedTime bson.M
		err = mongoDb.Collection("flowers").FindOneAndUpdate(ctx, filter, updateTime).Decode(&updatedTime)

		if err != nil {
			return nil, err
		}
	}
	return &ret, nil
}

func (mDb MongoDatabase) ModifyFlower(ctx context.Context, id ObjectID, newFlower Flower) (*Flower, error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":       newFlower.Name,
			"latin_name": newFlower.LatinName,
			"quantity":   newFlower.Quantity,
		},
	}

	if _, err := mongoDb.Collection("flowers").UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	}

	createdRecord := mongoDb.Collection("flowers").FindOne(ctx, filter)

	updatedFlower := &Flower{}
	if err := createdRecord.Decode(updatedFlower); err != nil {
		return nil, err
	}

	return updatedFlower, nil
}

func (mDb MongoDatabase) DeleteMultipleFlowers(ctx context.Context, flowerIDs []ObjectID) error {
	filter := bson.M{"_id": bson.M{"$in": flowerIDs}}
	_, err := mongoDb.Collection("flowers").DeleteMany(ctx, filter)
	return err
}

func (mDb MongoDatabase) UpdateVisibilityByTime(ctx context.Context, timestamp time.Time) (modified int64, err error) {
	filter := bson.M{"added_time": bson.M{"$lte": timestamp}}
	update := bson.M{"$set": bson.M{"visible": false}}
	updateResult, err := mongoDb.Collection("flowers").UpdateMany(ctx, filter, update)
	modified = updateResult.ModifiedCount
	return modified, err
}
