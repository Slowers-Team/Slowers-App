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

//? Expand Note to Notes (or a map)
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
	site := new(Site)

	if err := c.BodyParser(site); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	log.Println("received:", site)

	if site.Name == "" {
		return c.Status(400).SendString("Site name cannot be empty")
	}

	newSite := Site{Name: site.Name, Note: site.Note, AddedTime: time.Now(),
		Parent: site.Parent, Flowers: make([]*primitive.ObjectID, 0), Owner: site.Owner}

	insertResult, err := sites.InsertOne(c.Context(), newSite)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	} else {
		log.Println("created site", insertResult.InsertedID)
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := sites.FindOne(c.Context(), filter)

	createdSite := &Site{}
	createdRecord.Decode(createdSite)

	return c.Status(201).JSON(createdSite)
}

func getRootSites(c *fiber.Ctx) error {
	cursor, err := sites.Find(c.Context(), bson.D{{"parent", nil}})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var foundSites []Site
	if err := cursor.All(c.Context(), &foundSites); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println(foundSites)

	return c.JSON(foundSites)
}

func getSite(c *fiber.Ctx) error {
	id := c.Params("id")
	siteID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	var resultSite bson.M

	filter := bson.M{"_id": siteID}
	idErr := sites.FindOne(c.Context(), filter).Decode(&resultSite)

	if idErr != nil {
		if errors.Is(idErr, mongo.ErrNoDocuments) {
			log.Println("tried to find site", siteID, "but it doesn't exist")
		}
		return c.Status(500).SendString(idErr.Error())
	}

	log.Println("found site:", resultSite)

	matchStage := bson.D{{"$match", bson.D{{"parent", siteID}}}}
	sortStage := bson.D{{"$sort", bson.D{{"name", 1}}}}
	unsetStage := bson.D{{"$unset", bson.A{"parent", "addedTime", "owner", "flowers", "added_time"}}}

	cursor, err := sites.Aggregate(c.Context(), mongo.Pipeline{matchStage, sortStage, unsetStage})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var subSites []bson.M

	if err = cursor.All(c.Context(), &subSites); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	log.Println("subsites:", subSites)

	result := bson.M{"site": resultSite, "subsites": subSites}
	log.Println("result:", result)

	return c.JSON(result)
}
