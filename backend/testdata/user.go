package testdata

import "github.com/Slowers-team/Slowers-App/database"

func GetUser() database.User {
	return database.User{
		ID:       database.NewID("66fd465c0011335cd891aea7"),
		Username: "testuser",
		Email:    "testuser@test.com",
		Password: "testpassword",
		Role:     "grower",
	}
}
