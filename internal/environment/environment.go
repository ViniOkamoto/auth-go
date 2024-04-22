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
	JWTKey      string
	x
}

func Init() error {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	databaseURL := getVariable("DATABASE_URL")

	port := getVariable("PORT")
	portInt, err := strconv.Atoi(*port)
	if err != nil {
		panic(err)
	}

	jwtKey := getVariable("JWT_KEY")

	Config = &EnvironmentConfig{
		DatabaseURL: *databaseURL,
		Port:        portInt,
		JWTKey:      *jwtKey,
	}

	return nil
}

func getVariable(key string) *string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Missing %s environment variable", key)
		panic(errors.New("missing environment variable"))
	}

	return &value
}
