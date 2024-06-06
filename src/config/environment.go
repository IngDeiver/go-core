package config

import (
	"os"

	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	"github.com/joho/godotenv"
)


var L = logger.Get()

func LoadEnv() {
	env := os.Getenv("APP_ENV")
    var err error

    switch env {
    case "production":
        err = godotenv.Load(".env.production")
    case "development":
        err = godotenv.Load(".env.development")
    default:
        err = godotenv.Load(".env.local")
    }

    if err != nil {
        L.Fatal().Msg("Error loading .env file")
    }

	if len(env) > 0 {
		L.Info().Msgf("Environment loaded: %s", env)
	}else {
		L.Info().Msg("Environment loaded: local")
	}
}