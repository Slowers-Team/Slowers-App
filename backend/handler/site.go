package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/database"
)

func AddSite(c *fiber.Ctx) error {
	site := new(database.Site)

	if err := c.BodyParser(site); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	log.Println("received:", site)

	if site.Name == "" {
		return c.Status(400).SendString("Site name cannot be empty")
	}

	newSite := database.Site{Name: site.Name, Note: site.Note, AddedTime: time.Now(),
		Parent: site.Parent, Flowers: make([]*database.ObjectID, 0), Owner: site.Owner}

	createdSite, err := db.AddSite(c.Context(), newSite)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	} else {
		log.Println("created site", createdSite.ID)
	}

	return c.Status(201).JSON(createdSite)
}

func GetRootSites(c *fiber.Ctx) error {
	foundSites, err := db.GetRootSites(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	log.Println(foundSites)

	return c.JSON(foundSites)
}

func GetSite(c *fiber.Ctx) error {
	id := c.Params("id")

	if !database.IsValidID(id) {
		return c.SendStatus(400)
	}

	result, err := db.GetSite(c.Context(), id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	log.Println("result:", result)

	return c.JSON(result)
}

func DeleteSite(c *fiber.Ctx) error {
	id := c.Params("id")
	if !database.IsValidID(id) {
		return c.SendStatus(400)
	}

	deleteResult, err := db.DeleteSite(c.Context(), id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(deleteResult)
}
