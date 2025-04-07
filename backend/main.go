package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Slowers-team/Slowers-App/application"
	"github.com/Slowers-team/Slowers-App/databases/mongo"
	"github.com/Slowers-team/Slowers-App/databases/sql"
	"github.com/Slowers-team/Slowers-App/handlers"
	"github.com/cloudinary/cloudinary-go"
)

func main() {
	secretKey, databaseURI, port, env, envUseSQL, SQLDatabaseURI, envProdEnv := GetEnvironmentVariables()
	useSQL, err := strconv.ParseBool(envUseSQL)
	if err != nil {
		log.Fatal(err)
	}

	prodEnv, err := strconv.ParseBool(envProdEnv)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Using production environment: ", prodEnv)

	cld, err := cloudinary.New()

	if err != nil {
		log.Fatal(err)
	}

	cld.Config.URL.Secure = true

	mongoDb := mongo.NewMongoDatabase(databaseURI)
	if env == "test" {
		if err := mongoDb.Connect("SlowersTest"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := mongoDb.Connect("Slowers"); err != nil {
			log.Fatal(err)
		}
	}

	sqlDb := sql.NewSQLDatabase(SQLDatabaseURI)
	if env == "test" {
		err := sqlDb.Connect("slowerstest", false, prodEnv)
		if err != nil && strings.Contains(err.Error(), "failed to connect to") {
			// Try connecting again with 10 second cooldown to give time for database creation
			time.Sleep(10 * time.Second)
			if err = sqlDb.Connect("slowerstest", false, prodEnv); err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatal(err)
		}
	} else {
		err := sqlDb.Connect("slowers", false, prodEnv)
		if err != nil && strings.Contains(err.Error(), "failed to connect to") {
			// Try connecting again with 10 second cooldown to give time for database creation
			time.Sleep(10 * time.Second)
			if err = sqlDb.Connect("slowers", false, prodEnv); err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatal(err)
		}
	}

	application.SetSecretKey(secretKey)
	application.SetEnv(env)
	handlers.SetSecretKey(secretKey)
	handlers.SetDatabases(mongoDb, sqlDb)
	handlers.SetCloudinary(cld)

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

	mongoDbErr := mongoDb.Disconnect()
	sqlDbErr := sqlDb.Disconnect()

	if appErr != nil {
		log.Fatal(appErr)
	}
	if mongoDbErr != nil {
		log.Fatal(mongoDbErr)
	}
	if sqlDbErr != nil {
		log.Fatal(sqlDbErr)
	}
}
