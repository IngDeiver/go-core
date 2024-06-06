package config

import (
	"context"
	"os"

	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var Logger = logger.Get()

func InitMongoDB() {
	uri := os.Getenv("MONGO_URI")
	if len(uri) == 0{
		Logger.Fatal().Msg("MONGO_URI environment is required")
	}
	var err error
	clientOptions := options.Client().ApplyURI(uri)
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		Logger.Fatal().Err(err).Str("error","Mongo error")
	}

	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		Logger.Fatal().Err(err).Str("error","Mongo error")
	}

	Logger.Info().Msg("MongoDB connected")
}

func GetCollection(collectionName string) *mongo.Collection {
	databaseName := os.Getenv("MONGO_DB") 
	if len(databaseName) == 0{
		Logger.Fatal().Msg("MONGO_DB environment is required")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}