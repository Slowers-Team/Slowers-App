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
	thumbnailID, err := database.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println("got ID:", thumbnailID)
	thumbnail, err := db.GetImageByID(c.Context(), thumbnailID, "thumbnails")
	log.Println(thumbnailID, " -> ", thumbnail, err)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	filepath := fmt.Sprintf("./images/%v.%v", thumbnailID.Hex(), thumbnail.FileFormat)
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

func FetchThumbnailsByEntity(c *fiber.Ctx) error {
	entityID := c.Params("entityID")

	images, err := db.GetImagesByEntity(c.Context(), entityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(images)
}
