package utils

import (
	"encoding/json"
	"log"

	"github.com/Slowers-team/Slowers-App/database"
)

func FlowersToJSON(flowers []database.Flower) string {
	flowersJSON, err := json.Marshal(flowers)
	if err != nil {
		log.Fatal(err)
	}
	return string(flowersJSON)
}

func IDToJSON(id string) string {
	return "{\"id\": \"" + id + "\"}"
}
