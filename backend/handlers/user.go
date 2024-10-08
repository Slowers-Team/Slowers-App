package handlers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/utils"
)

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var SecretKey []byte

func SetSecretKey(newSecretKey []byte) {
	SecretKey = newSecretKey
}

func CreateUser(c *fiber.Ctx) error {
	user := new(database.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	log.Println(user)

	if user.Username == "" || user.Password == "" || user.Email == "" || user.Role == "" {
		return c.Status(400).SendString("All fields are required")
	}

	count, err := db.CountUsersWithEmail(c.Context(), user.Email)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if count > 0 {
		return c.Status(400).SendString("email already exists")
	}

	if !utils.IsEmailValid(user.Email) {
		return c.Status(400).SendString("invalid email")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newUser := database.User{Username: user.Username, Password: hashedPassword, Email: user.Email, Role: user.Role}

	err = db.CreateUser(c.Context(), newUser)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(201)
}

func HandleLogin(c *fiber.Ctx) error {
	login := new(LogIn)

	if err := c.BodyParser(login); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	user, err := db.GetUserByEmail(c.Context(), login.Email)
	if err != nil {
		return c.Status(401).SendString("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.Status(401).SendString("Invalid email or password")
	}

	claims := &jwt.StandardClaims{
		Subject:   primitive.ObjectID(user.ID).Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return c.Status(500).SendString("Could not create token")
	}

	return c.JSON(fiber.Map{"token": tokenString, "role": user.Role})
}
