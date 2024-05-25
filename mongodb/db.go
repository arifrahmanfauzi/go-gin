package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	client *mongo.Client
	db     string
}

func NewDB(ctx context.Context, client *mongo.Client, database string) *DB {
	return &DB{client: client, db: database}
}

func (d *DB) Collection() {

}
