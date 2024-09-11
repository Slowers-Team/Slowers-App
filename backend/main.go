package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Flower struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	LatinName string             `json:"latin_name" bson:"latin_name"`
	AddedTime time.Time          `json:"added_time" bson:"added_time"`
}

var collection *mongo.Collection

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
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

	app := fiber.New()

	app.Post("/api/flowers", addFlower)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	app.Static("/", "./client/dist")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func addFlower(c *fiber.Ctx) error {
	flower := new(Flower)

	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if flower.Name == "" {
		return c.Status(400).SendString("Flower name cannot be empty")
	}

	newFlower := Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now()}

	insertResult, err := collection.InsertOne(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdFlower := &Flower{}
	createdRecord.Decode(createdFlower)

	return c.Status(201).JSON(createdFlower)
}
