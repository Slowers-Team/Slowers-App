package main

import (
	"log"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/database"
	psqldatabase "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/handlers"
)

func main() {
	secretKey, databaseURI, port, env, useSQL, SQLDatabaseURI := GetEnvironmentVariables()

	db := database.NewMongoDatabase(databaseURI)
	if env == "test" {
		if err := db.Connect("SlowersTest"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := db.Connect("Slowers"); err != nil {
			log.Fatal(err)
		}
	}

	var sqldb *psqldatabase.SQLDatabase
	if useSQL == "true" {
		sqldb := psqldatabase.NewSQLDatabase(SQLDatabaseURI)
		if env == "test" {
			if err := sqldb.Connect("slowerstest"); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := sqldb.Connect("slowers"); err != nil {
				log.Fatal(err)
			}
		}
	}

	application.SetSecretKey(secretKey)
	application.SetEnv(env)
	handlers.SetSecretKey(secretKey)
	handlers.SetDatabase(db)

	app := application.SetupAndSetAuthTo(true)

	appErr := app.Listen("0.0.0.0:" + port)

	dbErr := db.Disconnect()

	if appErr != nil {
		log.Fatal(appErr)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	if useSQL == "true" {
		sqldb.Disconnect()
	}
}
