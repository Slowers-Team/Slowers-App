package handlers

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	image := new(database.Image)
	if err := c.BodyParser(image); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if image.Note == "" {
		return c.Status(400).SendString("Image note cannot be empty")
	}

	if image.Entity == nil || *image.Entity == database.NilObjectID {
		return c.Status(400).SendString("Entity associated to image cannot be null")
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	fileext := ""
	mimetype := file.Header["Content-Type"][0]
	if mimetype == "image/jpeg" {
		fileext = "jpg"
	} else if mimetype == "image/png" {
		fileext = "png"
	}

	if fileext == "" {
		return c.Status(400).SendString("Image should be in JPEG or PNG format")
	}

	if file.Size > 10485760 {
		return c.Status(400).SendString("Image cannot be larger than 10 MB")
	}

	if fileinfo, err := os.Stat("./images"); errors.Is(err, os.ErrNotExist) || !fileinfo.IsDir() {
		os.Remove("./images")
		if err := os.Mkdir("./images", 0775); err != nil {
			return c.Status(500).SendString("Could not create directory for images: " + err.Error())
		}
	}

	newImage := database.Image{FileFormat: fileext, Note: image.Note, Entity: image.Entity, Owner: userID}

	createdImage, err := db.AddImage(c.Context(), newImage)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	savepath := "./images/" + createdImage.ID.Hex() + "." + fileext
	if err := c.SaveFile(file, savepath); err != nil {
		db.DeleteImage(c.Context(), createdImage.ID)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(createdImage)
}

func DownloadImage(c *fiber.Ctx) error {
	filepath := "./images/" + c.Params("filename")

	if _, err := os.Stat(filepath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.SendStatus(404)
		} else {
			return c.Status(500).SendString(err.Error())
		}
	}

	return c.SendFile(filepath)
}

func DeleteImage(c *fiber.Ctx) error {
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

func FetchImagesByEntity(c *fiber.Ctx) error {
	entityID := c.Params("entityID")

	images, err := db.GetImagesByEntity(c.Context(), entityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(images)
}

func SetFavorite(c *fiber.Ctx) error {
	entityID := c.Params("entityID")
	entityType := c.Params("entityType")
	imageID := c.Params("imageID")

	log.Printf(entityID, entityType, imageID)

	return c.JSON(true)
}
