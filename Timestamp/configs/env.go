package configs

import (
    "log"
    "os"
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
    env := os.Getenv("APP_ENV")
	if env == "development" {
		env = "dev"
	} else if env == "production" {
		env = "prod"
	}

	filename := fmt.Sprintf(".env.%s", env)

	if err := godotenv.Load(filename); err != nil {
		log.Fatalf("Error loading %s file", filename)
	}
}