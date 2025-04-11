package testutils

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/databases/sql"
)

func ConnectMongoDB() mongo.Database {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	mongoDatabaseURI := os.Getenv("MONGODB_URI")
	if mongoDatabaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	MongoDb := mongo.NewMongoDatabase(mongoDatabaseURI)
	if err := MongoDb.Connect("SlowersTest"); err != nil {
		log.Fatal(err)
	}

	return MongoDb
}

func ConnectSqlDB() sql.Database {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	SQLDatabaseURI := os.Getenv("SQLDATABASEURI")
	if SQLDatabaseURI == "" {
		log.Fatal("Set your 'SQLDATABASEURI' environment variable.")
	}

	sqlDb := sql.NewSQLDatabase(SQLDatabaseURI)
	err := sqlDb.Connect("slowerstest", true, false)
	if err != nil && strings.Contains(err.Error(), "failed to connect to") {
		// Try connecting again with 10 second cooldown to give time for database creation
		time.Sleep(10 * time.Second)
		if err = sqlDb.Connect("slowerstest", true, false); err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	}

	return sqlDb
}

func DisconnectMongoDB(MongoDb mongo.Database) {
	err := MongoDb.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}

func DisconnectSqlDB(sqlDb sql.Database) {
	err := sqlDb.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
