package handlers

import (
	"fmt"

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
		return c.Status(400).SendString("Image name cannot be empty")
	}

	if image.Entity == nil {
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

	newImage := database.Image{Note: image.Note, Entity: image.Entity, Owner: userID}

	createdImage, err := db.AddImage(c.Context(), newImage)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	savepath := fmt.Sprintf("./images/%v.%v", createdImage.ID.Hex(), fileext)
	if err := c.SaveFile(file, savepath); err != nil {
		db.DeleteImage(c.Context(), createdImage.ID)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(createdImage)
}
