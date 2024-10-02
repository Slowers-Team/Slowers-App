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

func AddFlower(c *fiber.Ctx) error {
	flower := new(database.Flower)

	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if flower.Name == "" {
		return c.Status(400).SendString("Flower name cannot be empty")
	}

	newFlower := database.Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now()}

	createdFlower, err := db.AddFlower(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
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
