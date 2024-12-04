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
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowers, err := db.GetUserFlowers(c.Context(), userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
}

func AddFlower(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

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

	site, err := db.GetSiteByID(c.Context(), *flower.Site)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if site == nil {
		return c.Status(404).SendString("Site not found")
	}

	if flower.Quantity < 0 {
		return c.Status(400).SendString("Flower quantity cannot be negative")
	}

	newFlower := database.Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now(),
		Grower: &userID, GrowerEmail: grower.Email, Site: &site.ID, SiteName: site.Name, Quantity: flower.Quantity, Visible: false}

	createdFlower, err := db.AddFlower(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowerID := createdFlower.ID
	err = db.AddFlowerToSite(c.Context(), site.ID, flowerID)
	if err != nil {
		return c.Status(500).SendString("Failed to update site with flower ID: " + err.Error())
	}

	return c.Status(201).JSON(createdFlower)
}

func DeleteFlower(c *fiber.Ctx) error {
	id, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
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
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	siteID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	flowers, err := db.GetAllFlowersRelatedToSite(c.Context(), siteID, userID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
}

func ToggleFlowerVisibility(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowerID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newValue, err := db.ToggleFlowerVisibility(c.Context(), userID, flowerID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(newValue)
}

func DeleteMultipleFlowers(c *fiber.Ctx) error {
	var flowerIDs []string
	if err := c.BodyParser(&flowerIDs); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var ids []database.ObjectID
	for _, idStr := range flowerIDs {
		id, err := database.ParseID(idStr)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		ids = append(ids, id)
	}

	if err := db.DeleteMultipleFlowers(c.Context(), ids); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
