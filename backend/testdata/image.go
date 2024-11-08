package testdata

import (
	"log"

	"github.com/Slowers-team/Slowers-App/database"
)

func GetImages() []database.Image {
	imageIDStrs := []string{
		"671a48439de32ca350317e26",
		"671a4b659de32ca350317e27",
	}
	imageIDs := []database.ObjectID{}

	for _, idStr := range imageIDStrs {
		curImageID, err := database.ParseID(idStr)
		if err != nil {
			log.Fatal(err)
		}
		imageIDs = append(imageIDs, curImageID)
	}

	flowerID, err := database.ParseID("842af389e234e768923475bc")
	if err != nil {
		log.Fatal(err)
	}

	siteID, err := database.ParseID("66f5027d6430d371f8636c3c")
	if err != nil {
		log.Fatal(err)
	}

	ownerID, err := database.ParseID("66fd465c0011335cd891aea7")
	if err != nil {
		log.Fatal(err)
	}

	return []database.Image{
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

func GetImagesForAdding() []database.Image {
	images := GetImages()
	for i, _ := range images {
		images[i].ID = database.NilObjectID
	}
	return images
}
