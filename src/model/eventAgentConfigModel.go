package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventAgentConfig struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	EventName string             `bson:"eventName"`
	Active    bool               `bson:"active"`
}
