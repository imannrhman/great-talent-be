package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"great-talent-be/exception"
	"time"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()

	defer cancel()

	//mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	//exception.PanicIfNeeded(err)
	//
	//mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	//exception.PanicIfNeeded(err)
	//
	//mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	//exception.PanicIfNeeded(err)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	option := options.Client().
		ApplyURI(configuration.Get("MONGO_URI")).
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, option)

	exception.PanicIfNeeded(err)
	database := client.Database(configuration.Get("MONGO_DATABASE"))

	println("connected")
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
