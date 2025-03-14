package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	database "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/utils"
)

var SecretKey []byte

func SetSecretKey(newSecretKey []byte) {
	SecretKey = newSecretKey
}

func CreateUser(c *fiber.Ctx) error {
	user := new(database.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return c.Status(400).SendString("All fields are required")
	}

	// count, err := db.CountUsersWithEmail(c.Context(), user.Email)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }

	// if count > 0 {
	// 	return c.Status(400).SendString("email already exists")
	// }

	if !utils.IsEmailValid(user.Email) {
		return c.Status(400).SendString("invalid email")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Println(timestamp)

	isactive := true

	isadmin := false

	newUser := database.User{
		CreatedAt:    user.CreatedAt,
		LastModified: user.LastModified,
		LastLogin:    timestamp,
		Username:     user.Username,
		Password:     hashedPassword,
		Email:        user.Email,
		IsActive:     isactive,
		IsAdmin:      isadmin}

	fmt.Println("Uusi käyttäjä:", newUser)

	createdUser, err := db.CreateUser(c.Context(), newUser)

	if err != nil {
		fmt.Println("Wörkkiikö/ollaan nyt errorissa")
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("Printti jos toimii")

	// otetaan printit pois sitten kun tämä on kokonaan valmis

	return LogUserIn(c, createdUser, 201)
}

// func HandleLogin(c *fiber.Ctx) error {
// 	login := new(database.LogIn)

// 	if err := c.BodyParser(login); err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	user, err := db.GetUserByEmail(c.Context(), login.Email)
// 	if err != nil {
// 		return c.Status(401).SendString("Invalid email or password")
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
// 	if err != nil {
// 		return c.Status(401).SendString("Invalid email or password")
// 	}

// 	return LogUserIn(c, user, 200)
// }

func LogUserIn(c *fiber.Ctx, user *database.User, status int) error {
	claims := &jwt.StandardClaims{
		Subject:   strconv.Itoa(user.ID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return c.Status(500).SendString("Could not create token")
	}

	return c.Status(status).JSON(fiber.Map{"token": tokenString, "username": user.Username})
}

// func GetUser(c *fiber.Ctx) error {
// 	userID, err := GetCurrentUser(c)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	result, err := db.GetUserByID(c.Context(), userID)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	return c.JSON(result)
// }

// func SetRole(c *fiber.Ctx) error {
// 	userID, err := GetCurrentUser(c)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	var role string
// 	if err := c.BodyParser(&role); err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	_, err = enums.RoleFromString(role)
// 	if err != nil {
// 		return c.Status(400).SendString(err.Error())
// 	}

// 	err = db.SetUserRole(c.Context(), userID, role)
// 	if err != nil {
// 		return c.Status(500).SendString(err.Error())
// 	}

// 	return c.Status(201).JSON(role)
// }
