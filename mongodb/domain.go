package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type querying interface {
	col() *mongo.Collection
	// build builds the filter structure with previous operations.
	build() bson.D
}
