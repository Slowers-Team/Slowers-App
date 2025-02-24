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
	thumbnail, err := db.GetImageByID(c.Context(), thumbnailID)
	log.Println(thumbnailID, " -> ", thumbnail, err)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	filepath := fmt.Sprintf("./thumbnails/%v.%v", thumbnailID.Hex(), thumbnail.FileFormat)
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
