package handlers

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetThumbnailByID(c *fiber.Ctx) error {
	imageID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println("got ID:", imageID)
	image, err := db.GetImageByID(c.Context(), imageID)
	log.Println(imageID, " -> ", image, err)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	filepath := fmt.Sprintf("./images/%v.%v", imageID.Hex(), image.FileFormat)
	log.Println(filepath)

	if _, err := os.Stat(filepath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("404 fail")
			return c.SendStatus(404)
		} else {
			log.Println("500 fail")
			return c.Status(500).SendString(err.Error())
		}
	}
	log.Println("sending file")

	return c.SendFile(filepath)
}

func DeleteThumbnail(c *fiber.Ctx) error {
	id, err := database.ParseID(c.Params("id"))
	log.Printf("Received ID for deletion: %s", id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid image ID format")
	}

	deleted, err := db.DeleteImage(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if !deleted {
		return c.Status(fiber.StatusNotFound).SendString("Image not found")
	}

	extensions := []string{"jpg", "png"}
	found := false

	for _, ext := range extensions {
		imagePath := fmt.Sprintf("./images/%s.%s", id.Hex(), ext)
		if _, err := os.Stat(imagePath); err == nil {
			if err := os.Remove(imagePath); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error deleting image file")
			}
			log.Printf("Successfully deleted image file: %s", imagePath)
			found = true
			break
		}
	}

	if !found {
		log.Printf("Image file not found for ID: %s", id.Hex())
		return c.Status(fiber.StatusNotFound).SendString("Image file not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func FetchThumbnailsByEntity(c *fiber.Ctx) error {
	entityID := c.Params("entityID")

	images, err := db.GetImagesByEntity(c.Context(), entityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(images)
}
