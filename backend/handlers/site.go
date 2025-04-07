package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
)

func AddSite(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	site := new(mongo.Site)

	if err := c.BodyParser(site); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if site.Name == "" {
		return c.Status(400).SendString("Site name cannot be empty")
	}

	var flowers []*mongo.ObjectID
	if site.Flowers != nil {
		flowers = site.Flowers
	} else {
		flowers = make([]*mongo.ObjectID, 0)
	}

	newSite := mongo.Site{Name: site.Name, Note: site.Note, AddedTime: time.Now(),
		Parent: site.Parent, Flowers: flowers, Owner: &userID}

	createdSite, err := MongoDb.AddSite(c.Context(), newSite)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	} else {
		log.Println("created site", createdSite.ID)
	}

	return c.Status(201).JSON(createdSite)
}

func GetRootSites(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	foundSites, err := MongoDb.GetRootSites(c.Context(), userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(foundSites)
}

func GetSite(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	siteID, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result, err := MongoDb.GetSite(c.Context(), siteID, userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(result)
}

func DeleteSite(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	siteID, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	deleteResult, err := MongoDb.DeleteSite(c.Context(), siteID, userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(deleteResult)
}
