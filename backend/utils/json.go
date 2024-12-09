package utils

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Slowers-team/Slowers-App/database"
)

func FlowerToJSON(flower database.Flower) []byte {
	flowerJSON, err := json.Marshal(flower)
	if err != nil {
		log.Fatal(err)
	}
	return flowerJSON
}

func FlowersToJSON(flowers []database.Flower) []byte {
	flowersJSON, err := json.Marshal(flowers)
	if err != nil {
		log.Fatal(err)
	}
	return flowersJSON
}

func SiteToJSON(site database.Site) []byte {
	siteJSON, err := json.Marshal(site)
	if err != nil {
		log.Fatal(err)
	}
	return siteJSON
}

func SitesToJSON(sites []database.Site) []byte {
	sitesJSON, err := json.Marshal(sites)
	if err != nil {
		log.Fatal(err)
	}
	return sitesJSON
}

func SiteDataToJSON(siteData bson.M) []byte {
	siteDataJSON, err := json.Marshal(siteData)
	if err != nil {
		log.Fatal(err)
	}
	return siteDataJSON
}

func UserToJSON(user database.User) []byte {
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return userJSON
}

func LogInToJSON(login database.LogIn) []byte {
	loginJSON, err := json.Marshal(login)
	if err != nil {
		log.Fatal(err)
	}
	return loginJSON
}

func ImageToJSON(image database.Image) []byte {
	imageJSON, err := json.Marshal(image)
	if err != nil {
		log.Fatal(err)
	}
	return imageJSON
}

func ImagesToJSON(images []database.Image) []byte {
	imagesJSON, err := json.Marshal(images)
	if err != nil {
		log.Fatal(err)
	}
	return imagesJSON
}

func IDToJSON(id string) string {
	return "{\"id\": \"" + id + "\"}"
}

func ToJSON(val any) []byte {
	asJSON, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}
	return asJSON
}
