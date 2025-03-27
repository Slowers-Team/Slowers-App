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
	"github.com/cloudinary/cloudinary-go"
)

func main() {
	secretKey, databaseURI, port, env, envUseSQL, SQLDatabaseURI := GetEnvironmentVariables()
	useSQL, err := strconv.ParseBool(envUseSQL)

	if err != nil {
		log.Fatal(err)
	}

	cld, err := cloudinary.New()

	if err != nil {
		log.Fatal(err)
	}

	cld.Config.URL.Secure = true

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
	handlers.SetCloudinary(cld)

	var sqldb *psqldatabase.SQLDatabase
	if useSQL {
		sqldb := psqldatabase.NewSQLDatabase(SQLDatabaseURI)
		if env == "test" {
			if err := sqldb.Connect("slowerstest", false); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := sqldb.Connect("slowers", false); err != nil {
				log.Fatal(err)
			}
		}

		psqlHandlers.SetSecretKey(secretKey) //TODO: Check if needed
		psqlHandlers.SetDatabase(*sqldb)
	}

	app := application.SetupAndSetAuthTo(true, useSQL)

	ticker := time.NewTicker(24 * time.Hour)
	if time.Now().Hour() == 0 {
		ticker = time.NewTicker(8 * time.Minute)
	}
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// log.Println("1 minute has passed")
				timestamp := time.Now().AddDate(0, -6, 0)
				modified, err := handlers.UpdateVisibilityByTime(timestamp)
				if err != nil {
					log.Println("0 modified, error:" + err.Error())
				}
				log.Println(strconv.Itoa(int(modified)) + " set invisible")
			case <-quit:
				ticker.Stop()
				return
			}
		}
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
