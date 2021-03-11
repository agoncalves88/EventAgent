package infrastructure

import (
	"context"
	"log"

	"EventAgent.Consumer/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

type MongoBase struct {
	Configuration configuration.Configuration
}

func (m MongoBase) GetDatabase() *mongo.Database {
	clientOptions := options.Client().ApplyURI(m.Configuration.MongoConnection)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("events")
}
