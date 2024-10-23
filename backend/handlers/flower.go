package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetFlowers(c *fiber.Ctx) error {
	flowers, err := db.GetFlowers(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
}

func GetUserFlowers(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	if !database.IsValidID(user) {
		return c.Status(500).SendString("Malformed userID in header")
	}
	userID := database.NewID(user)

	flowers, err := db.GetUserFlowers(c.Context(), userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
}

func GetFlowerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if !database.IsValidID(id) {
		return c.SendStatus(400)
	}

	flowerID := database.NewID(id)

	flower, err := db.GetFlowerByID(c.Context(), flowerID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flower)
}

func AddFlower(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	if !database.IsValidID(user) {
		return c.Status(500).SendString("Malformed userID in header")
	}
	userID := database.NewID(user)

	grower, err := db.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(500).SendString("User not found: " + err.Error())
	}

	flower := new(database.Flower)
	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if flower.Name == "" {
		return c.Status(400).SendString("Flower name cannot be empty")
	}

	if flower.Site == nil {
		return c.Status(400).SendString("SiteID is required")
	}
	if !database.IsValidID(flower.Site.Hex()) {
		return c.Status(400).SendString("Invalid siteID")
	}
	siteID := database.NewID(flower.Site.Hex())

	site, err := db.GetSiteByID(c.Context(), siteID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if site == nil {
		return c.Status(404).SendString("Site not found")
	}

	newFlower := database.Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now(), Grower: &userID, GrowerEmail: grower.Email, Site: &siteID, SiteName: site.Name}

	createdFlower, err := db.AddFlower(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowerID := createdFlower.ID
	err = db.AddFlowerToSite(c.Context(), siteID, flowerID)
	if err != nil {
		return c.Status(500).SendString("Failed to update site with flower ID: " + err.Error())
	}

	return c.Status(201).JSON(createdFlower)
}

func DeleteFlower(c *fiber.Ctx) error {
	id := c.Params("id")
	if !database.IsValidID(id) {
		return c.SendStatus(400)
	}

	anyDeleted, err := db.DeleteFlower(c.Context(), id)
	if err != nil {
		return c.SendStatus(500)
	}
	if !anyDeleted {
		return c.SendStatus(404)
	}

	return c.SendStatus(204)
}

func GetSiteFlowers(c *fiber.Ctx) error {
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	if !database.IsValidID(user) {
		return c.Status(500).SendString("Malformed userID in header")
	}
	userID := database.NewID(user)

	siteID := c.Params("id")

	flowers, err := db.GetAllFlowersRelatedToSite(c.Context(), siteID, userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
}
