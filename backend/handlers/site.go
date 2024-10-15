package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/database"
)

func AddSite(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	if !database.IsValidID(user) {
		return c.Status(500).SendString("Malformed userID in header")
	}
	userID := database.NewID(user)

	site := new(database.Site)

	if err := c.BodyParser(site); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if site.Name == "" {
		return c.Status(400).SendString("Site name cannot be empty")
	}

	var flowers []*database.ObjectID
	if site.Flowers != nil {
		flowers = site.Flowers
	} else {
		flowers = make([]*database.ObjectID, 0)
	}

	newSite := database.Site{Name: site.Name, Note: site.Note, AddedTime: time.Now(),
		Parent: site.Parent, Flowers: flowers, Owner: &userID}

	createdSite, err := db.AddSite(c.Context(), newSite)
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

	foundSites, err := db.GetRootSites(c.Context(), userID)
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

	siteID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	result, err := db.GetSite(c.Context(), siteID, userID)
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

	siteID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	deleteResult, err := db.DeleteSite(c.Context(), siteID, userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(deleteResult)
}
