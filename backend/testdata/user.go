package testdata

import (
	"log"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetUser() database.User {
	userID, err := database.ParseID("66fd465c0011335cd891aea7")
	if err != nil {
		log.Fatal(err)
	}
	return database.User{
		ID:       userID,
		Username: "testuser",
		Email:    "testuser@test.com",
		Password: "testpassword",
		Role:     "grower",
	}
}
