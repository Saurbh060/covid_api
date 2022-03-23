package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoFields struct {
	State      string  `bson:"state"`
	TotalCases float64 `bson:"totalCases"`
}

type StateData struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StateCases []MongoFields      `bson:"stateCases" json:"stateCases"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
