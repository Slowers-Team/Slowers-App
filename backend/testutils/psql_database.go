package testutils

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	database "github.com/Slowers-team/Slowers-App/database/psql"
)

func ConnectPsqlDB() database.Database {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	databaseURI := os.Getenv("SQLTESTDATABASEURI")
	if databaseURI == "" {
		log.Fatal("Set your 'SQLTESTDATABASEURI' environment variable.")
	}

	db := database.NewSQLDatabase(databaseURI)
	if err := db.Connect("SlowersTest"); err != nil {
		log.Fatal(err)
	}

	return db
}

func DisconnectPsqlDB(db database.Database) {
	err := db.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
