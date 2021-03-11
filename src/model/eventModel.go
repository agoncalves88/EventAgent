package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Payload     string             `bson:"payload"`
	Completed   bool               `bson:"completed"`
	UrlCallback string             `bson:"urlCallback"`
}
