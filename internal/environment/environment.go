package environment

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *EnvironmentConfig

type EnvironmentConfig struct {
	DatabaseURL string
	Port        int
}

func Init() error {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		return err
	}

	databaseURL, err := getVariable("DATABASE_URL")

	if err != nil {
		log.Fatal(err)
		return err
	}
	port, err := getVariable("PORT")

	if err != nil {
		log.Fatal(err)
		return err
	}
	portInt, err := strconv.Atoi(*port)

	if err != nil {
		log.Fatal(err)
		return err
	}

	Config = &EnvironmentConfig{
		DatabaseURL: *databaseURL,
		Port:        portInt,
	}

	return nil
}

func getVariable(key string) (*string, error) {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Missing %s environment variable", key)
		return nil, errors.New("missing environment variable")
	}

	return &value, nil
}
