package testdata

import (
	"log"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetUsers() []database.User {
	userIDStrs := []string{
		"66fd465c0011335cd891aea7",
		"670ea95dc96af530f69341d5",
		"67a61d4b6d544410878b2edc",
		"67a61d676d544410878b2edd",
		"67a61d8d6d544410878b2ede",
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
		{
			ID:       userIDs[2],
			Username: "testuser3",
			Email:    "testuser3@test.com",
			Password: "testpassword",
			Role:     "retailer",
		},
		{
			ID:       userIDs[3],
			Username: "testuser4",
			Email:    "testuser4@test.com",
			Password: "testpassword",
			Role:     "growerowner",
		},
		{
			ID:       userIDs[4],
			Username: "testuser5",
			Email:    "testuser5@test.com",
			Password: "testpassword",
			Role:     "retailerowner",
		},
	}
}
