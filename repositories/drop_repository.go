package repositories

import (
	"context"
	"go-starter/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DropRepository struct {
	client *mongo.Client
}

func NewDropRepository(client *mongo.Client) *DropRepository {
	return &DropRepository{client: client}
}

func (r *DropRepository) GetAllDrops(ctx context.Context) ([]models.Drops, error) {
	collection := r.client.Database("db_superkul_order").Collection("customerTripPlanningDt")
	var drops []models.Drops

	cur, err := collection.Find(ctx, struct{}{}, options.Find().SetLimit(20))
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var drop models.Drops
		err := cur.Decode(&drop)
		if err != nil {
			return nil, err
		}
		drops = append(drops, drop)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return drops, nil
}

func (r *DropRepository) FetchDrops() {
	//pipeline := mongodb.Pipeline{}

}
