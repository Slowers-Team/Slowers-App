package testdata

import (
	"log"
	"time"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetTestFlowers() []database.Flower {
	flowerIDStrs := []string{
		"842af389e234e768923475bc",
		"485a28e70545c378ff29b438",
		"ac83264ff67837e87eb82322",
	}
	flowerIDs := []database.ObjectID{}

	for _, idStr := range flowerIDStrs {
		curFlowerID, err := database.ParseID(idStr)
		if err != nil {
			log.Fatal(err)
		}
		flowerIDs = append(flowerIDs, curFlowerID)
	}

	growerID := GetUsers()[0].ID
	siteID, err := database.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}

	return []database.Flower{
		{
			ID:        flowerIDs[0],
			Name:      "sunflower",
			LatinName: "Helianthus annuus",
			AddedTime: time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
			Grower:    &growerID,
			Site:      &siteID,
			Visible:   true,
		},
		{
			ID:        flowerIDs[1],
			Name:      "daisy",
			LatinName: "Bellis perennis",
			AddedTime: time.Date(2024, 8, 27, 7, 4, 32, 0, time.UTC),
			Grower:    &growerID,
			Site:      &siteID,
			Visible:   true,
		},
		{
			ID:        flowerIDs[2],
			Name:      "spring crocus",
			LatinName: "Crocus vernus",
			AddedTime: time.Date(2024, 9, 29, 19, 58, 1, 0, time.UTC),
			Grower:    &growerID,
			Site:      &siteID,
			Visible:   true,
		},
	}
}

func GetTestFlowerForUser2() database.Flower {
	flowerID, err := database.ParseID("66fd466f0011335cd891aea8")
	if err != nil {
		log.Fatal(err)
	}

	growerID := GetUsers()[1].ID
	siteID, err := database.ParseID("6700042668d22894f711af60")
	if err != nil {
		log.Fatal(err)
	}

	return database.Flower{
		ID:        flowerID,
		Name:      "cornflower",
		LatinName: "Centaurea cyanus",
		AddedTime: time.Date(2024, 10, 11, 19, 32, 17, 0, time.UTC),
		Grower:    &growerID,
		Site:      &siteID,
		Visible:   true,
	}
}
