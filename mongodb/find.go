package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type Find struct {
	query querying
}

func NewFind(collection *mongo.Collection) *Find {
	return &Find{}
}
