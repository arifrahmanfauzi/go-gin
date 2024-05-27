package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Drops struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	DispatchNumber string             `bson:"dispatchNumber" json:"dispatch_number"`
	Job            string
}
