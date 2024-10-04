package testdata

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetRootSites() []database.Site {
	flowerID := database.NewID("842af389e234e768923475bc")
	ownerID := database.NewID("66fd465c0011335cd891aea7")
	return []database.Site{
		{
			ID: database.NewID("66f5027d6430d371f8636c3c"),
			Name: "Greenhouse A",
			AddedTime: time.Date(2024, 9, 26, 6, 43, 9, 0, time.UTC),
			Note: "Just a note",
			Parent: nil,
			Flowers: []*database.ObjectID{&flowerID},
			Owner: &ownerID,
		},
	}
}

func GetSite() bson.M {
	return bson.M{
		"site": GetRootSites()[0],
		"subsites": []bson.M{
			{
				"_id": "66ffd761a00aa71bdad57262",
				"name": "Pot 1",
				"note": "This is another note",
			},
		},
	}
}
