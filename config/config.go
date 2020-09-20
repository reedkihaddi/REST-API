package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/reedkihaddi/REST-API/logging"
)

// InitConfig loads the env variables.
func InitConfig() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "local"
	}
	// godotenv.Load(".env." + env)
	err := godotenv.Load("./config/" + env + ".env")
	if err != nil {
		logging.Log.Fatalf("error loading .env file err:%s", err)
	}
}
