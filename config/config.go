package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	AppStage string
	Port     string
	DBUser   string
	DBPass   string
	DBSchema string
	DBHost   string
	DBPort   string
}

func ReadConfig() (c Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// vars without default value
	required := []string{
		"ENVIRONMENT",
		"PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_SCHEMA",
	}

	missing := []string{}
	for _, reqVar := range required {
		_, present := os.LookupEnv(reqVar)
		if !present {
			missing = append(missing, reqVar)
		}
	}

	if len(missing) > 0 {
		return c, fmt.Errorf("missing the following ENV variables: %s", strings.Join(missing, ", "))
	}

	c = Config{
		AppStage: os.Getenv("ENVIRONMENT"),
		Port:     os.Getenv("PORT"),
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASSWORD"),
		DBSchema: os.Getenv("DB_SCHEMA"),
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
	}

	return c, nil

}