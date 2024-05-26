package helpers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pipeline struct {
	stages []bson.D
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
func (p *Pipeline) Match(filters ...bson.E) *Pipeline {
	match := bson.D{}

	if len(filters) > 0 {
		match = bson.D{{Key: "$match",
			Value: bson.D{{Key: "$and", Value: bson.A{filters}}}}}
	}

	p.stages = append(p.stages, match)
	return p
}

func (p *Pipeline) Sort(sortField string, sortOrder int) *Pipeline {
	sort := bson.D{
		{Key: "$sort", Value: bson.D{{Key: sortField, Value: sortOrder}}},
	}

	p.stages = append(p.stages, sort)
	return p
}

func (p *Pipeline) Limit(limit int) *Pipeline {
	sort := bson.D{
		{Key: "$limit", Value: limit},
	}

	p.stages = append(p.stages, sort)
	return p
}

func (p *Pipeline) LookUp(from, localField, foreignField, as string) *Pipeline {
	lookup := bson.D{
		{Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: from},
				{Key: "localField", Value: localField},
				{Key: "foreignField", Value: foreignField},
				{Key: "as", Value: as},
			},
		},
	}

	p.stages = append(p.stages, lookup)
	return p
}
func (p *Pipeline) UnwindStage(path string, preserveNullAndEmptyArrays bool) *Pipeline {
	unwind := bson.D{
		{Key: "$unwind",
			Value: bson.D{
				{Key: "path", Value: path},
				{Key: "preserveNullAndEmptyArrays", Value: preserveNullAndEmptyArrays},
			},
		},
	}

	p.stages = append(p.stages, unwind)
	return p
}
func (p *Pipeline) ProjectStage(projections ...bson.E) *Pipeline {
	project := bson.D{}

	if len(projections) > 0 {
		project = bson.D{{Key: "$project", Value: bson.A{projections}}}
	}

	p.stages = append(p.stages, project)
	return p
}

// Add other stage methods as needed...

func (p *Pipeline) Build() mongo.Pipeline {
	return p.stages
}
