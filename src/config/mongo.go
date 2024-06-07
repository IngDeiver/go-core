package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client


func InitMongoDB() {
	uri := os.Getenv("MONGO_URI")
	if len(uri) == 0{
		l.Fatal().Msg("MONGO_URI environment is required")
	}
	var err error
	clientOptions := options.Client().ApplyURI(uri)
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	if err == nil{
		l.Info().Msg("MongoDB connected")
	}

}

func GetCollection(collectionName string) *mongo.Collection {
	databaseName := os.Getenv("MONGO_DB") 
	if len(databaseName) == 0{
		l.Fatal().Msg("MONGO_DB environment is required")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}