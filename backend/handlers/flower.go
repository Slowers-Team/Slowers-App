package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
)

func ValidateFlower(flower mongo.Flower) error {
	if flower.Name == "" {
		return errors.New("Flower name cannot be empty")
	}
	if flower.Site == nil {
		return errors.New("SiteID is required")
	}
	if flower.Quantity < 0 {
		return errors.New("Flower quantity cannot be negative")
	}
	return nil
}

func GetFlowers(c *fiber.Ctx) error {
	flowers, err := MongoDb.GetFlowers(c.Context())
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

	flowers, err := MongoDb.GetUserFlowers(c.Context(), userID)
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

	grower, err := sqlDb.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(500).SendString("User not found: " + err.Error())
	}

	flower := new(mongo.Flower)
	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = ValidateFlower(*flower)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	site, err := MongoDb.GetSiteByID(c.Context(), *flower.Site)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if site == nil {
		return c.Status(404).SendString("Site not found")
	}

	newFlower := mongo.Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now(),
		Grower: &userID, GrowerEmail: grower.Email, Site: &site.ID, SiteName: site.Name, Quantity: flower.Quantity, Visible: false}

	createdFlower, err := MongoDb.AddFlower(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowerID := createdFlower.ID
	err = MongoDb.AddFlowerToSite(c.Context(), site.ID, flowerID)
	if err != nil {
		return c.Status(500).SendString("Failed to update site with flower ID: " + err.Error())
	}

	return c.Status(201).JSON(createdFlower)
}

func DeleteFlower(c *fiber.Ctx) error {
	id, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	anyDeleted, err := MongoDb.DeleteFlower(c.Context(), id)
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

	siteID, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	flowers, err := MongoDb.GetAllFlowersRelatedToSite(c.Context(), siteID, userID)
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

	flowerID, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newValue, err := MongoDb.ToggleFlowerVisibility(c.Context(), userID, flowerID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(newValue)
}

func ModifyFlower(c *fiber.Ctx) error {
	id, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	flower := new(mongo.Flower)
	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = ValidateFlower(*flower)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	updatedFlower, err := MongoDb.ModifyFlower(c.Context(), id, mongo.Flower{Name: flower.Name, LatinName: flower.LatinName, Quantity: flower.Quantity})

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(updatedFlower)
}

func DeleteMultipleFlowers(c *fiber.Ctx) error {
	var flowerIDs []string
	if err := c.BodyParser(&flowerIDs); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var ids []mongo.ObjectID
	for _, idStr := range flowerIDs {
		id, err := mongo.ParseID(idStr)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		ids = append(ids, id)
	}

	if err := MongoDb.DeleteMultipleFlowers(c.Context(), ids); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

func UpdateVisibilityByTime(timestamp time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	modified, err := mongo.Database.UpdateVisibilityByTime(MongoDb, ctx, timestamp)
	if err != nil {
		return 0, err
	}
	return modified, err
}
