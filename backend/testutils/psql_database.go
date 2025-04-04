package testutils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
	err := sqldb.Connect("slowerstest", true, false)
	if err != nil && strings.Contains(err.Error(), "failed to connect to") {
		// Try connecting again with 10 second cooldown to give time for database creation
		time.Sleep(10 * time.Second)
		if err = sqldb.Connect("slowerstest", true, false); err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
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
