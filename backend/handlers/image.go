package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strings"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/utils"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

func ValidateImage(image mongo.Image) error {
	if image.Note == "" {
		return errors.New("Image note cannot be empty")
	}
	if image.Entity == nil || *image.Entity == mongo.NilObjectID {
		return errors.New("Entity associated to image cannot be null")
	}
	return nil
}

func ValidateFile(file multipart.FileHeader) error {
	if file.Size > 10485760 {
		return errors.New("Image cannot be larger than 10 MB")
	}
	if file.Size <= 0 {
		return errors.New("Image size cannot be zero or negative")
	}
	return nil
}

func UploadImage(c *fiber.Ctx) error {
	userID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	image := new(mongo.Image)
	if err := c.BodyParser(image); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = ValidateImage(*image)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	fileext, err := utils.SetImageFormat(file.Header["Content-Type"][0])
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err = ValidateFile(*file)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newImage := mongo.Image{FileFormat: fileext, Note: image.Note, Entity: image.Entity, Owner: userID}

	createdImage, err := MongoDb.AddImage(c.Context(), newImage)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Read the file into a byte slice
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Create an io.Reader from the byte slice
	fileReader := bytes.NewReader(fileBytes)

	// THIS PART IS IF YOU WANT TO BRING BACK THUMBNAILS, IT DOES NOT WORK RIGHT NOW
	// if filedir, err := os.Stat("./thumbnails"); errors.Is(err, os.ErrNotExist) || !filedir.IsDir() {
	// 	os.Remove("./thumbnails")
	// 	if err := os.Mkdir("./thumbnails", 0775); err != nil {
	// 		return c.Status(500).SendString("Could not create directory for thumbnails: " + err.Error())
	// 	}
	// }

	// savepath = "./thumbnails/" + createdImage.ID.Hex() + "." + fileext
	// createdThumbnail, err := os.Create(savepath)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }

	resp, err := cld.Upload.Upload(c.Context(), fileReader, uploader.UploadParams{
		PublicID:       "images/" + createdImage.ID.Hex(),
		UniqueFilename: true,
		Overwrite:      true,
	})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println(resp)

	// err = utils.ResizeImage(fileReader, createdThumbnail, fileext, 200, 200)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }

	return c.Status(201).JSON(createdImage)
}

func GetImageByID(c *fiber.Ctx) error {
	imageID, err := mongo.ParseID(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println("got ID:", imageID.Hex())
	image, err := MongoDb.GetImageByID(c.Context(), imageID)
	log.Println(imageID, " -> ", image, err)
	if err != nil {
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	// filepath := fmt.Sprintf("./images/%v.%v", imageID.Hex(), image.FileFormat)
	// log.Println(filepath)

	resp, err := cld.Admin.Asset(c.Context(), admin.AssetParams{PublicID: "images/" + imageID.Hex()})
	if err != nil {
		fmt.Println("error")
	}

	// if _, err := os.Stat(filepath); err != nil {
	// 	if errors.Is(err, os.ErrNotExist) {
	// 		log.Println("404 fail")
	// 		return c.SendStatus(404)
	// 	} else {
	// 		log.Println("500 fail")
	// 		return c.Status(500).SendString(err.Error())
	// 	}
	// }
	log.Println("sending file", resp.SecureURL)

	// return c.SendFile(filepath)

	return c.SendString(resp.SecureURL)
}

func DownloadImage(c *fiber.Ctx) error {
	filename := c.Params("filename")
	image := strings.Split(filename, ".")[0]

	resp, err := cld.Admin.Asset(c.Context(), admin.AssetParams{PublicID: "images/" + image})
	if err != nil {
		fmt.Println("error")
	}

	return c.SendString(resp.SecureURL)
}

func DeleteImage(c *fiber.Ctx) error {
	id, err := mongo.ParseID(c.Params("id"))
	log.Printf("Received ID for deletion: %s", id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid image ID format")
	}

	deletedImage, err := MongoDb.DeleteImage(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if !deletedImage {
		return c.Status(fiber.StatusNotFound).SendString("Image not found")
	}

	// This is yet to be implemented, commented to not make unnecessary errors
	// deletedThumbnail, err := db.DeleteImage(c.Context(), id, "thumbnails")
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	// }
	// if !deletedThumbnail {
	// 	return c.Status(fiber.StatusNotFound).SendString("Image not found")
	// }

	// extensions := []string{"jpg", "png"}
	// found := false

	// if !found {
	// 	log.Printf("Image file not found for ID: %s", id.Hex())
	// 	return c.Status(fiber.StatusNotFound).SendString("Image file not found")
	// }

	resp, err := cld.Upload.Destroy(c.Context(), uploader.DestroyParams{
		PublicID:     "images/" + id.Hex(),
		ResourceType: "image"})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting image file from image server")
	}

	_ = resp

	return c.SendStatus(fiber.StatusNoContent)
}

func FetchImagesByEntity(c *fiber.Ctx) error {
	entityID := c.Params("entityID")

	images, err := MongoDb.GetImagesByEntity(c.Context(), entityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(images)
}

func SetFavorite(c *fiber.Ctx) error {
	type favoriteData struct {
		EntityID   string `json:"entityID"`
		EntityType string `json:"entityType"`
		ImageID    string `json:"imageID"`
	}

	formData := new(favoriteData)
	if err := c.BodyParser(formData); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	EntityID, err := mongo.ParseID(formData.EntityID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Invalid entity ID format: %v", formData.EntityID))
	}

	collection, err := utils.SetCollectionType(formData.EntityType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	ImageID, err := mongo.ParseID(formData.ImageID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Invalid image ID format: %v", formData.ImageID))
	}

	UserID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Could not get current user")
	}

	err = MongoDb.SetFavoriteImage(c.Context(), UserID, EntityID, ImageID, collection)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).SendString("")
}

func ClearFavorite(c *fiber.Ctx) error {
	type favoriteData struct {
		EntityID   string `json:"entityID"`
		EntityType string `json:"entityType"`
	}

	formData := new(favoriteData)
	if err := c.BodyParser(formData); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	EntityID, err := mongo.ParseID(formData.EntityID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Invalid entity ID format: %v", formData.EntityID))
	}

	collection, err := utils.SetCollectionType(formData.EntityType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	UserID, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Could not get current user")
	}

	err = MongoDb.ClearFavoriteImage(c.Context(), UserID, EntityID, collection)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).SendString("")
}
