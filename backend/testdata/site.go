package testdata

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
)

func GetRootSites() []mongo.Site {
	siteID, err := mongo.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}
	flowerID := GetFlowers()[0].ID
	ownerID := string(GetUsers()[0].ID)
	return []mongo.Site{
		{
			ID:        siteID,
			Name:      "Greenhouse A",
			AddedTime: time.Date(2024, 9, 26, 6, 43, 9, 0, time.UTC),
			Note:      "Just a note",
			Parent:    nil,
			Flowers:   []*mongo.ObjectID{&flowerID},
			Owner:     &ownerID,
		},
	}
}

func GetRootSitesForUser2() []mongo.Site {
	siteID, err := mongo.ParseID("6700042668d22894f711af60")
	if err != nil {
		log.Fatal(err)
	}
	flowerID := GetFlowerForUser2().ID
	ownerID := string(GetUsers()[1].ID)
	return []mongo.Site{
		{
			ID:        siteID,
			Name:      "Field 1",
			AddedTime: time.Date(2024, 9, 19, 12, 8, 49, 0, time.UTC),
			Note:      "",
			Parent:    nil,
			Flowers:   []*mongo.ObjectID{&flowerID},
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

func PrepareSiteForAdding(site mongo.Site) mongo.Site {
	site.ID = mongo.NilObjectID
	site.FavoriteImage = ""
	return site
}
