package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"EventAgent.Consumer/configuration"
	"EventAgent.Consumer/infrastructure"
	"EventAgent.Consumer/model"
	"go.mongodb.org/mongo-driver/bson"
)

var _eventCollection = "EventAgentConfig"

func main() {
	env := flag.String("env", "", "")
	flag.Parse()
	config := configuration.GetConfig(string(*env))

	db := infrastructure.MongoBase{Configuration: config}.GetDatabase()
	kafkaBase := infrastructure.KafkaBase{Configuration: config}

	ctx := context.Background()
	query := &bson.M{
		"active": true,
	}
	cur, err := db.Collection(_eventCollection).Find(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	events := []string{}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		result := model.EventAgentConfig{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, result.EventName)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(events))

	for _, event := range events {
		go func(topic string) {
			kafkaBase.Consumer(ctx, topic)
			waitGroup.Done()
		}(event)
	}

	waitGroup.Wait()
}
