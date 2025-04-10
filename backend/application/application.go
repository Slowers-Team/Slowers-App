package application

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/Slowers-team/Slowers-App/handlers"
	"github.com/Slowers-team/Slowers-App/testdata"
)

var SecretKey []byte
var Env string

func SetSecretKey(newSecretKey []byte) {
	SecretKey = newSecretKey
}

func SetEnv(newEnv string) {
	Env = newEnv
}

func SetupAndSetAuthTo(isAuthOn bool) *fiber.App {
	app := fiber.New()

	app.Get("/api/healthcheck", handlers.HealthCheck)
	app.Post("/api/register", handlers.CreateUser)
	app.Post("/api/login", handlers.HandleLogin)

	if Env == "test" {
		app.Get("/api/reset", handlers.ResetDatabase)
	}

	api := app.Group("/api")

	if isAuthOn {
		api.Use(AuthMiddleware)
	} else {
		api.Use(TestAuthMiddleware)
	}

	api.Post("/flowers", handlers.AddFlower)
	api.Get("/flowers", handlers.GetFlowers)
	api.Get("/flowers/user", handlers.GetUserFlowers)
	api.Delete("/flowers/:id", handlers.DeleteFlower)
	api.Post("/flowers/:id/visibility", handlers.ToggleFlowerVisibility)
	api.Put("/flowers/:id", handlers.ModifyFlower)
	api.Post("/flowers/delete-multiple", handlers.DeleteMultipleFlowers)

	api.Post("/sites", handlers.AddSite)
	api.Get("/sites", handlers.GetRootSites)
	api.Get("/sites/:id", handlers.GetSite)
	api.Delete("/sites/:id", handlers.DeleteSite)
	api.Get("/sites/:id/flowers", handlers.GetSiteFlowers)

	api.Get("/user", handlers.GetUser)

	api.Get("/user/designation", handlers.GetDesignation)

	api.Post("/business", handlers.CreateBusiness)
	api.Get("/business", handlers.GetBusiness)

	api.Post("/membership", handlers.AddMembership)
	api.Get("/membership/all", handlers.GetAllMembersInBusiness)
	api.Delete("/membership", handlers.DeleteMembership)

	api.Post("/images", handlers.UploadImage)
	api.Get("/images/id/:id", handlers.GetImageByID)
	api.Get("/images/:filename", handlers.DownloadImage)
	api.Get("/images/entity/:entityID", handlers.FetchImagesByEntity)
	api.Delete("/images/:id", handlers.DeleteImage)
	api.Post("/images/favorite", handlers.SetFavorite)
	api.Post("/images/clearfavorite", handlers.ClearFavorite)

	api.Get("/thumbnails/id/:id", handlers.GetThumbnailByID)

	app.Static("/assets", "./client/dist/assets")

	app.Static("/*", "./client/dist")
	return app
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := &jwt.RegisteredClaims{}
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
	c.Locals("userID", testdata.GetUsers()[0].ID)
	return c.Next()
}
