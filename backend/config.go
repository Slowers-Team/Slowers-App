package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariables() ([]byte, string, string, string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	envSecretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(envSecretKey) == 0 {
		log.Fatal("Set your SECRET_KEY as an environment variable.")
	}

	envDatabaseURI := os.Getenv("MONGODB_URI")
	if envDatabaseURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = "5001"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	envSQLDatabaseURI := os.Getenv("SQLDATABASEURI")
	if envSQLDatabaseURI == "" {
		log.Fatal("Set your 'SQLDATABASEURI' environment variable or do not enable SQL-connection.")
	}

	envProdEnv := os.Getenv("PRODUCTION_ENVIRONMENT")
	if envProdEnv == "" {
		envProdEnv = "false"
	}

	return envSecretKey, envDatabaseURI, envPort, env, envSQLDatabaseURI, envProdEnv
}
