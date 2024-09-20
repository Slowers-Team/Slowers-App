package main

import (
	"context"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Flower struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	LatinName string             `json:"latin_name" bson:"latin_name"`
	AddedTime time.Time          `json:"added_time" bson:"added_time"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

var collection *mongo.Collection
var userCollection *mongo.Collection

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
	userCollection = client.Database("Slowers").Collection("users")

	app := fiber.New()

	app.Post("/api/flowers", addFlower)
	app.Get("/api/flowers", getFlowers)
	app.Delete("/api/flowers/:id", deleteFlower)
	app.Post("/api/register", createUser)
	app.Get("/api/users", getUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	app.Static("/", "./client/dist")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getFlowers(c *fiber.Ctx) error {
	cursor, err := collection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowers := make([]Flower, 0)
	if err := cursor.All(c.Context(), &flowers); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(flowers)
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

func deleteFlower(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	filter := bson.M{"_id": objectID}
	result, err := collection.DeleteOne(c.Context(), filter)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.SendStatus(204)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(400).SendString("All fields are required")
	}

	count, err := userCollection.CountDocuments(c.Context(), bson.M{"email": user.Email})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if count > 0 {
		return c.Status(400).SendString("email already exists")
	}

	if !isEmailValid(user.Email) {
		return c.Status(400).SendString("invalid email")
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newUser := User{Username: user.Username, Password: hashedPassword, Email: user.Email}

	insertResult, err := userCollection.InsertOne(c.Context(), newUser)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := userCollection.FindOne(c.Context(), filter)

	createdUser := &User{}
	createdRecord.Decode(createdUser)

	return c.Status(201).JSON(createdUser)
}

func getUsers(c *fiber.Ctx) error {
	cursor, err := userCollection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []User
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
