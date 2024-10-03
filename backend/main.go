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
	user, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).SendString("Invalid userID in header")
	}
	userID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flower := new(Flower)
	if err := c.BodyParser(flower); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if flower.Name == "" {
		return c.Status(400).SendString("Flower name cannot be empty")
	}

	if flower.Site == nil {
		return c.Status(400).SendString("SiteID is required")
	}

	siteID, err := primitive.ObjectIDFromHex(flower.Site.Hex())
	if err != nil {
		return c.Status(400).SendString("Invalid siteID")
	}

	newFlower := Flower{Name: flower.Name, LatinName: flower.LatinName, AddedTime: time.Now(), Grower: &userID, Site: &siteID}

	insertResult, err := collection.InsertOne(c.Context(), newFlower)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	flowerID := insertResult.InsertedID.(primitive.ObjectID)

	update := bson.M{"$push": bson.M{"flowers": flowerID}}
	_, err = sites.UpdateOne(c.Context(), bson.M{"_id": siteID}, update)
	if err != nil {
		return c.Status(500).SendString("Failed to update site with flower ID: " + err.Error())
	}

	filter := bson.M{"_id": insertResult.InsertedID}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdFlower := &Flower{}
	err = createdRecord.Decode(createdFlower)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

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
	err = createdRecord.Decode(createdUser)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
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

func handleLogin(c *fiber.Ctx) error {

	login := new(LogIn)

	if err := c.BodyParser(login); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	user := new(User)

	err := userCollection.FindOne(c.Context(), bson.D{{Key: "email", Value: login.Email}}).Decode(&user)
	if err != nil {
		return c.Status(401).SendString("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.Status(401).SendString("Invalid email or password")
	}

	claims := &jwt.StandardClaims{
		Subject:   user.ID.Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return c.Status(500).SendString("Could not create token")
	}

	return c.JSON(fiber.Map{"token": tokenString})
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
