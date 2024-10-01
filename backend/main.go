package main

import (
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/handler"
)

var SecretKey []byte

func main() {
	var databaseURI, port string
	SecretKey, databaseURI, port = GetEnvironmentVariables()

	databaseClient, err := database.Connect(databaseURI, "Slowers")
	if err != nil {
		log.Fatal(err)
	}

	handler.SetSecretKey(SecretKey)
	handler.SetDatabase(new(database.ActualDatabase))

	app := Setup()
	app.Static("/", "./client/dist")

	appErr := app.Listen("0.0.0.0:" + port)

	dbErr := database.Disconnect(databaseClient)

	if appErr != nil {
		log.Fatal(appErr)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}

func Setup() *fiber.App {
	app := fiber.New()

	app.Post("/api/register", handler.CreateUser)
	app.Post("/api/login", handler.HandleLogin)

	app.Post("/api/sites", handler.AddSite)
	app.Get("/api/sites", handler.GetRootSites)
	app.Get("/api/sites/:id", handler.GetSite)
	app.Delete("/api/sites/:id", handler.DeleteSite)

	//app.Use(AuthMiddleware)

	app.Post("/api/flowers", handler.AddFlower)
	app.Get("/api/flowers", handler.GetFlowers)
	app.Delete("/api/flowers/:id", handler.DeleteFlower)

	return app
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	c.Locals("userID", claims.Subject)
	return c.Next()
}
