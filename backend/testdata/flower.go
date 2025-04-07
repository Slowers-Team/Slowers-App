package testdata

import (
	"log"
	"time"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
)

func GetFlowers() []mongo.Flower {
	flowerIDStrs := []string{
		"842af389e234e768923475bc",
		"485a28e70545c378ff29b438",
		"ac83264ff67837e87eb82322",
	}
	flowerIDs := []mongo.ObjectID{}

	for _, idStr := range flowerIDStrs {
		curFlowerID, err := MongoDb.ParseID(idStr)
		if err != nil {
			log.Fatal(err)
		}
		flowerIDs = append(flowerIDs, curFlowerID)
	}

	grower := GetUsers()[0]
	growerID := grower.ID
	growerEmail := grower.Email

	siteID, err := MongoDb.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}
	siteName := "Greenhouse A"

	return []mongo.Flower{
		{
			ID:          flowerIDs[0],
			Name:        "sunflower",
			LatinName:   "Helianthus annuus",
			AddedTime:   time.Date(2024, 9, 30, 21, 11, 54, 0, time.UTC),
			Grower:      &growerID,
			GrowerEmail: growerEmail,
			Site:        &siteID,
			SiteName:    siteName,
			Quantity:    389,
			Visible:     true,
		},
		{
			ID:          flowerIDs[1],
			Name:        "daisy",
			LatinName:   "Bellis perennis",
			AddedTime:   time.Date(2024, 8, 27, 7, 4, 32, 0, time.UTC),
			Grower:      &growerID,
			GrowerEmail: growerEmail,
			Site:        &siteID,
			SiteName:    siteName,
			Quantity:    1,
			Visible:     true,
		},
		{
			ID:          flowerIDs[2],
			Name:        "spring crocus",
			LatinName:   "Crocus vernus",
			AddedTime:   time.Date(2024, 9, 29, 19, 58, 1, 0, time.UTC),
			Grower:      &growerID,
			GrowerEmail: growerEmail,
			Site:        &siteID,
			SiteName:    siteName,
			Quantity:    0,
			Visible:     true,
		},
	}
}

func PrepareFlowerForAdding(flower mongo.Flower) mongo.Flower {
	flower.ID = mongo.NilObjectID
	flower.FavoriteImage = ""
	return flower
}

func GetFlowerForUser2() mongo.Flower {
	flowerID, err := mongo.ParseID("66fd466f0011335cd891aea8")
	if err != nil {
		log.Fatal(err)
	}

	grower := GetUsers()[1]
	growerID := grower.ID
	growerEmail := grower.Email

	siteID, err := mongo.ParseID("6700042668d22894f711af60")
	if err != nil {
		log.Fatal(err)
	}
	siteName := "Field 1"

	return mongo.Flower{
		ID:          flowerID,
		Name:        "cornflower",
		LatinName:   "Centaurea cyanus",
		AddedTime:   time.Date(2024, 10, 11, 19, 32, 17, 0, time.UTC),
		Grower:      &growerID,
		GrowerEmail: growerEmail,
		Site:        &siteID,
		SiteName:    siteName,
		Quantity:    7,
		Visible:     true,
	}
}
