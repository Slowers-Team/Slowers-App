package main

import (
	"log"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/database"
)

func main() {
	secretKey, databaseURI, port := GetEnvironmentVariables()

	db := database.NewMongoDatabase(databaseURI)
	if err := db.Connect("Slowers"); err != nil {
		log.Fatal(err)
	}

	app := application.SetupAndSetAuthTo(true)
	app.Static("/", "./client/dist")

	appErr := app.Listen("0.0.0.0:" + port)

	dbErr := db.Disconnect()

	if appErr != nil {
		log.Fatal(appErr)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}
}
