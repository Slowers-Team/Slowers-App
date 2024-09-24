package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//? Expand Note to Notes (or a map)
//? SubSites as []ID or []*Site
type Site struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name"`
	AddedTime time.Time           `json:"added_time" bson:"added_time"`
	Note      string              `json:"note"`
	Parent    *primitive.ObjectID `json:"parent"`
	Flowers   []Flower            `json:"flowers"`
	Owner     *primitive.ObjectID `json:"owner"`
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
		Parent: site.Parent, Flowers: make([]Flower, 0), Owner: site.Owner}

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
