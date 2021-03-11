package main

import (
	"context"
	"flag"
	"log"

	"EventAgent.Consumer/configuration"
	"EventAgent.Consumer/infrastructure"
	"EventAgent.Consumer/model"
	"go.mongodb.org/mongo-driver/bson"
)

var _eventCollection = "Event"

func main() {
	env := flag.String("env", "", "")
	flag.Parse()
	config := configuration.GetConfig(string(*env))

	db := infrastructure.MongoBase{Configuration: config}.GetDatabase()

	ctx := context.Background()
	query := &bson.M{
		"completed": false,
	}
	cur, err := db.Collection(_eventCollection).Find(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		result := model.EventModel{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
