package testdata

import (
	"log"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
)

func GetImages() []mongo.Image {
	imageIDStrs := []string{
		"671a48439de32ca350317e26",
		"671a4b659de32ca350317e27",
	}
	imageIDs := []mongo.ObjectID{}

	for _, idStr := range imageIDStrs {
		curImageID, err := mongo.ParseID(idStr)
		if err != nil {
			log.Fatal(err)
		}
		imageIDs = append(imageIDs, curImageID)
	}

	flowerID, err := mongo.ParseID("842af389e234e768923475bc")
	if err != nil {
		log.Fatal(err)
	}

	siteID, err := mongo.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}

	ownerID := "1"

	return []mongo.Image{
		{
			ID:         imageIDs[0],
			FileFormat: "jpg",
			Note:       "Picture of sunflower",
			Entity:     &flowerID,
			Owner:      ownerID,
		},
		{
			ID:         imageIDs[1],
			FileFormat: "png",
			Note:       "Picture of greenhouse A",
			Entity:     &siteID,
			Owner:      ownerID,
		},
	}
}

func PrepareImageForAdding(image mongo.Image) mongo.Image {
	image.ID = mongo.NilObjectID
	return image
}
