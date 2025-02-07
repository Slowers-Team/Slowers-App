package testdata

import (
	"log"

	"github.com/Slowers-team/Slowers-App/database"
	"github.com/Slowers-team/Slowers-App/utils"
)

func GetUsers() []database.User {
	userIDStrs := []string{
		"66fd465c0011335cd891aea7",
		"670ea95dc96af530f69341d5",
	}
	userIDs := []database.ObjectID{}

	for _, idStr := range userIDStrs {
		curUserID, err := database.ParseID(idStr)
		if err != nil {
			log.Fatal(err)
		}
		userIDs = append(userIDs, curUserID)
	}

	return []database.User{
		{
			ID:       userIDs[0],
			Username: "testuser",
			Email:    "testuser@test.com",
			Password: "testpassword",
			Role:     "grower",
		},
		{
			ID:       userIDs[1],
			Username: "testuser2",
			Email:    "testuser2@test.com",
			Password: "testpassword2",
			Role:     "grower",
		},
	}
}

func PrepareUserForAdding(user database.User) database.User {
	user.ID = database.NilObjectID
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	return user
}
