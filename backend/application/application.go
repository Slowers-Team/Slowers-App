package application

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/Slowers-team/Slowers-App/handlers"
	"github.com/Slowers-team/Slowers-App/testdata"
)

var SecretKey []byte

func SetSecretKey(newSecretKey []byte) {
	SecretKey = newSecretKey
}

func SetupAndSetAuthTo(isAuthOn bool) *fiber.App {
	app := fiber.New()

	app.Post("/api/register", handlers.CreateUser)
	app.Post("/api/login", handlers.HandleLogin)

	app.Static("/assets", "./client/dist/assets")

	if isAuthOn {
		app.Use(AuthMiddleware)
	} else {
		app.Use(TestAuthMiddleware)
	}

	app.Post("/api/flowers", handlers.AddFlower)
	app.Get("/api/flowers", handlers.GetFlowers)
	app.Get("/api/flowers/user", handlers.GetUserFlowers)
	app.Delete("/api/flowers/:id", handlers.DeleteFlower)
	app.Post("/api/flowers/:id/visibility", handlers.ToggleFlowerVisibility)
	app.Put("/api/flowers/:id", handlers.ModifyFlower)

	app.Post("/api/sites", handlers.AddSite)
	app.Get("/api/sites", handlers.GetRootSites)
	app.Get("/api/sites/:id", handlers.GetSite)
	app.Delete("/api/sites/:id", handlers.DeleteSite)
	app.Get("/api/sites/:id/flowers", handlers.GetSiteFlowers)

	app.Get("/api/user", handlers.GetUser)
	app.Post("/api/user/role", handlers.SetRole)

	app.Post("/api/images", handlers.UploadImage)
	app.Get("/api/images/:filename", handlers.DownloadImage)
	app.Get("/api/images/entity/:entityID", handlers.FetchImagesByEntity)
	app.Delete("/api/images/:id", handlers.DeleteImage)

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

func TestAuthMiddleware(c *fiber.Ctx) error {
	c.Locals("userID", testdata.GetUsers()[0].ID.Hex())
	return c.Next()
}
