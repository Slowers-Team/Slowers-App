package application

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/handlers"
)

var SecretKey []byte

func SetSecretKey(newSecretKey []byte) {
	SecretKey = newSecretKey
}

func SetupAndSetAuthTo(isAuthOn bool) *fiber.App {
	app := fiber.New()

	app.Post("/api/register", handlers.CreateUser)
	app.Post("/api/login", handlers.HandleLogin)

	app.Post("/api/sites", handlers.AddSite)
	app.Get("/api/sites", handlers.GetRootSites)
	app.Get("/api/sites/:id", handlers.GetSite)
	app.Delete("/api/sites/:id", handlers.DeleteSite)

	if (isAuthOn) {
		app.Use(AuthMiddleware)
	}

	app.Post("/api/flowers", handlers.AddFlower)
	app.Get("/api/flowers", handlers.GetFlowers)
	app.Delete("/api/flowers/:id", handlers.DeleteFlower)

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
