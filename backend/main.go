package main

import (
	"log"
	"strconv"
	"time"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/database"
	psqldatabase "github.com/Slowers-team/Slowers-App/database/psql"
	"github.com/Slowers-team/Slowers-App/handlers"
	psqlHandlers "github.com/Slowers-team/Slowers-App/handlersPsql"
	"github.com/Slowers-team/Slowers-App/utils"
)

func main() {
	secretKey, databaseURI, port, env, envUseSQL, SQLDatabaseURI := GetEnvironmentVariables()
	useSQL, err := strconv.ParseBool(envUseSQL)

	if err != nil {
		log.Fatal(err)
	}

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

	application.SetSecretKey(secretKey)
	application.SetEnv(env)
	handlers.SetSecretKey(secretKey)
	handlers.SetDatabase(db)

	var sqldb *psqldatabase.SQLDatabase
	if useSQL {
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

		psqlHandlers.SetSecretKey(secretKey) //TODO: Check if needed
		psqlHandlers.SetDatabase(*sqldb)
	}

	app := application.SetupAndSetAuthTo(true, useSQL)

	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		utils.VisibilityTicker(ticker, quit)
	}()

	appErr := app.Listen("0.0.0.0:" + port)

	dbErr := db.Disconnect()

	if appErr != nil {
		log.Fatal(appErr)
	}
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	if useSQL {
		sqldb.Disconnect()
	}
}
