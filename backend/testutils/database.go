package testutils

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Slowers-team/Slowers-App/database"
)

func ConnectDB() database.Database {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	databaseURI := os.Getenv("MONGODB_URI")
	if databaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	db := database.NewMongoDatabase(databaseURI)
	if err := db.Connect("SlowersTest"); err != nil {
		log.Fatal(err)
	}

	return db
}

func DisconnectDB(db database.Database) {
	err := db.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
