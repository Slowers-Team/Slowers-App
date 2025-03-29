package testutils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	database "github.com/Slowers-team/Slowers-App/database/psql"
)

func ConnectSQLDB() database.Database {
	fmt.Println(os.Getwd())
	if err := godotenv.Load("../../../.dev.env"); err != nil {
		log.Println("No .env file found")
	}

	SQLDatabaseURI := os.Getenv("SQLDATABASEURI")
	if SQLDatabaseURI == "" {
		log.Fatal("Set your 'SQLDATABASEURI' environment variable or do not enable SQL-connection.")
	}

	sqldb := database.NewSQLDatabase(SQLDatabaseURI)
	if err := sqldb.Connect("slowerstest", true, false); err != nil {
		log.Fatal(err)
	}

	return sqldb
}

func DisconnectSQLDB(db database.Database) {
	err := db.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
