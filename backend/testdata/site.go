package testdata

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetRootSites() []database.Site {
	siteID, err := database.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}
	flowerID := GetTestFlowers()[0].ID
	ownerID := GetUsers()[0].ID
	return []database.Site{
		{
			ID:        siteID,
			Name:      "Greenhouse A",
			AddedTime: time.Date(2024, 9, 26, 6, 43, 9, 0, time.UTC),
			Note:      "Just a note",
			Parent:    nil,
			Flowers:   []*database.ObjectID{&flowerID},
			Owner:     &ownerID,
		},
	}
}

func GetRootSitesForUser2() []database.Site {
	siteID, err := database.ParseID("6700042668d22894f711af60")
	if err != nil {
		log.Fatal(err)
	}
	flowerID := GetTestFlowerForUser2().ID
	ownerID := GetUsers()[1].ID
	return []database.Site{
		{
			ID:        siteID,
			Name:      "Field 1",
			AddedTime: time.Date(2024, 9, 19, 12, 8, 49, 0, time.UTC),
			Note:      "",
			Parent:    nil,
			Flowers:   []*database.ObjectID{&flowerID},
			Owner:     &ownerID,
		},
	}
}

func GetSite() bson.M {
	return bson.M{
		"site": GetRootSites()[0],
		"subsites": []bson.M{
			{
				"_id":  "66ffd761a00aa71bdad57262",
				"name": "Pot 1",
				"note": "This is another note",
			},
		},
	}
}
