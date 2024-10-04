package utils

import (
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Slowers-team/Slowers-App/database"
)

func FlowerToJSON(flower database.Flower) string {
	flowerJSON, err := json.Marshal(flower)
	if err != nil {
		log.Fatal(err)
	}
	return string(flowerJSON)
}

func FlowersToJSON(flowers []database.Flower) string {
	flowersJSON, err := json.Marshal(flowers)
	if err != nil {
		log.Fatal(err)
	}
	return string(flowersJSON)
}

func SiteToJSON(site database.Site) string {
	siteJSON, err := json.Marshal(site)
	if err != nil {
		log.Fatal(err)
	}
	return string(siteJSON)
}

func SitesToJSON(sites []database.Site) string {
	sitesJSON, err := json.Marshal(sites)
	if err != nil {
		log.Fatal(err)
	}
	return string(sitesJSON)
}

func SiteDataToJSON(siteData bson.M) string {
	siteDataJSON, err := json.Marshal(siteData)
	if err != nil {
		log.Fatal(err)
	}
	return string(siteDataJSON)
}

func UserToJSON(user database.User) string {
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return string(userJSON)
}

func LogInToJSON(login database.LogIn) string {
	loginJSON, err := json.Marshal(login)
	if err != nil {
		log.Fatal(err)
	}
	return string(loginJSON)
}

func IDToJSON(id string) string {
	return "{\"id\": \"" + id + "\"}"
}
