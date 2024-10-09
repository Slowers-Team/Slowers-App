package main

import (
	"context"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Flower struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name"`
	LatinName string              `json:"latin_name" bson:"latin_name"`
	AddedTime time.Time           `json:"added_time" bson:"added_time"`
	Grower    *primitive.ObjectID `json:"grower"`
	Site      *primitive.ObjectID `json:"site"`
	SiteName  string              `json:"site_name" bson:"site_name"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var collection *mongo.Collection
var userCollection *mongo.Collection
var sites *mongo.Collection

var SecretKey []byte

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	SecretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(SecretKey) == 0 {
		log.Fatal("Set your SECRET_KEY as an environment variable.")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	clientOptions := options.Client().ApplyURI(mongoURI).SetTimeout(10 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB")

	collection = client.Database("Slowers").Collection("flowers")
	userCollection = client.Database("Slowers").Collection("users")
	sites = client.Database("Slowers").Collection("sites")

	app := fiber.New()

	app.Post("/api/register", createUser)
	app.Post("/api/login", handleLogin)

	app.Use(AuthMiddleware)

	app.Post("/api/flowers", addFlower)
	app.Get("/api/flowers", getFlowers)
	app.Get("/api/flowers/user", getUserFlowers)
	app.Delete("/api/flowers/:id", deleteFlower)

	app.Post("/api/sites", addSite)
	app.Get("/api/sites", getRootSites)
	app.Get("/api/sites/:id", getSite)
	app.Delete("/api/sites/:id", deleteSite)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	app.Static("/", "./client/dist")

	appErr := app.Listen("0.0.0.0:" + port)

	dbErr := db.Disconnect()

	if appErr != nil {
		log.Fatal(appErr)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}