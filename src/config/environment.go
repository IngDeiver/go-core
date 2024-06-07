package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func LoadEnv() {
	env := os.Getenv("APP_ENV")
    var err error

    switch env {
    case "production":
        err = godotenv.Load(".env.production")
        gin.SetMode(gin.ReleaseMode)
    case "development":
        err = godotenv.Load(".env.development")
    default:
        err = godotenv.Load(".env.local")
    }

    if err != nil {
        l.Fatal().Msg("Error loading .env file")
    }

	if len(env) > 0 {
		l.Info().Msgf("Environment loaded: %s", env)
	}else {
		l.Info().Msg("Environment loaded: local")
	}
}