package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`
	Todo_description string             `json:"todo_description,omitempty" bson:"todo_description,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedDate      time.Time          `json:"createdDate"`
	// LastUpdate       primitive.Timestamp `json:"lastUpdate"`
}
