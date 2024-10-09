package main

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ? Expand Note to Notes (or a map)
type Site struct {
	ID        primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string                `json:"name"`
	AddedTime time.Time             `json:"added_time" bson:"added_time"`
	Note      string                `json:"note"`
	Parent    *primitive.ObjectID   `json:"parent"`
	Flowers   []*primitive.ObjectID `json:"flowers"`
	Owner     *primitive.ObjectID   `json:"owner"`
}

func addSite(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	site := new(Site)

	if err := c.BodyParser(site); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if site.Name == "" {
		return c.Status(400).SendString("Site name cannot be empty")
	}

	var flowers []*primitive.ObjectID
	if site.Flowers != nil {
		flowers = site.Flowers
	} else {
		flowers = make([]*primitive.ObjectID, 0)
	}

	newSite := Site{Name: site.Name, Note: site.Note, AddedTime: time.Now(),
		Parent: site.Parent, Flowers: flowers, Owner: &userID}

	insertResult, err := sites.InsertOne(c.Context(), newSite)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	} else {
		log.Println("created site", insertResult.InsertedID)
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := sites.FindOne(c.Context(), filter)

	createdSite := &Site{}
	err = createdRecord.Decode(createdSite)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(createdSite)
}

func getRootSites(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	cursor, err := sites.Find(c.Context(), bson.M{"parent": nil, "owner": userID})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var foundSites []Site
	if err := cursor.All(c.Context(), &foundSites); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(foundSites)
}

func getSite(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id := c.Params("id")
	siteID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	var resultSite bson.M

	filter := bson.M{"_id": siteID, "owner": userID}
	idErr := sites.FindOne(c.Context(), filter).Decode(&resultSite)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find site", siteID, "but it doesn't exist")
		}
		return c.Status(500).SendString(idErr.Error())
	}

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "parent", Value: siteID}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "name", Value: 1}}}}
	unsetStage := bson.D{{Key: "$unset", Value: bson.A{"parent", "addedTime", "owner", "flowers", "added_time"}}}

	cursor, err := sites.Aggregate(c.Context(), mongo.Pipeline{matchStage, sortStage, unsetStage})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var subSites []bson.M

	if err = cursor.All(c.Context(), &subSites); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	result := bson.M{"site": resultSite, "subsites": subSites}

	return c.JSON(result)
}

func deleteSite(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {pipeline
		return c.Status(500).SendString("Invalid userID in header")
	}
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id := c.Params("id")
	siteID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
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
	// Open up array of documents to a stream of documents
	unwindStage := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$id"},
		},
		}}
	// Strip down everything except _id for each child Site
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "id", Value: "$related._id"},
		}},
	}

	cursor, err := sites.Aggregate(c.Context(), mongo.Pipeline{matchStage, graphLookupStage, projectStage, unwindStage})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var result []bson.M
	if err := cursor.All(c.Context(), &result); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	ids := make([]primitive.ObjectID, 0)
	ids = append(ids, siteID)

	for _, res := range result {
		sub_id, err := res["id"].(primitive.ObjectID)
		if !err {
			return c.Status(500).SendString("Fetched sub site ID was of wrong type")
		}
		ids = append(ids, sub_id)
	}

	log.Println("DELETE sites", ids)

	flowerDeleteFilter := bson.M{"site": bson.M{"$in": ids}}
	deleteFlowerResult, err := collection.DeleteMany(c.Context(), flowerDeleteFilter)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println("Deleted flowers count:", deleteFlowerResult.DeletedCount)

	deleteFilter := bson.M{"_id": bson.M{"$in": ids}}
	deleteResult, err := sites.DeleteMany(c.Context(), deleteFilter)
	if err != nil {
		log.Println("DELETE FAILED: ", err)
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(deleteResult)
}
