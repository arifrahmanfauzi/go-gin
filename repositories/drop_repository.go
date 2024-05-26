package repositories

import (
	"context"
	"go-gin/helpers"
	"go-gin/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *DropRepository) FetchDrops(ctx context.Context) (*[]bson.M, error) {
	pipeline := helpers.NewPipeline().
		LookUp("customerTripPlanning", "tripIdObject", "_id", "cTP").Sort("created_at", -1).
		UnwindStage("$cTP", true).Limit(20).ProjectStage(bson.E{Key: "_id", Value: 0},
		bson.E{Key: "dropId", Value: "$_id"},
		bson.E{Key: "dispatchNumber", Value: 1},
		bson.E{Key: "tripNumber", Value: "$cTP.tripNumber"},
		bson.E{Key: "created_at", Value: 1}).Build()
	// Run the aggregation
	collection := r.client.Database("db_superkul_order").Collection("customerTripPlanningDt")
	cursor, err := collection.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		log.Fatal("error aggregate : ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Process the results
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	return &results, nil
}
