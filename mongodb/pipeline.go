package mongodb

import "go.mongodb.org/mongo-driver/bson"

type Pipeline struct {
	stages []bson.D
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Match(filters ...bson.E) *Pipeline {
	match := bson.D{}
	if len(filters) > 0 {
		match = bson.D{{"$match",
			bson.D{{"$and", bson.A{filters}}}}}
	}
	p.stages = append(p.stages, match)
	return p
}
func (p *Pipeline) Or() {

}
func (p *Pipeline) Limit() {

}
func (p *Pipeline) Sort(field string, order int) *Pipeline {
	sort := bson.D{
		{"$sort", bson.D{{field, order}}},
	}

	p.stages = append(p.stages, sort)
	return p
}
func (p *Pipeline) Lookup(from, local, foreign, as string) *Pipeline {
	lookup := bson.D{
		{"$lookup",
			bson.D{
				{"from", from},
				{"localField", local},
				{"foreignField", foreign},
				{"as", as},
			},
		},
	}

	p.stages = append(p.stages, lookup)
	return p
}
func (p *Pipeline) Unwind(path string, preserveNullAndEmptyArrays bool) bson.D {
	return bson.D{
		{"$unwind",
			bson.D{
				{"path", path},
				{"preserveNullAndEmptyArrays", preserveNullAndEmptyArrays},
			},
		},
	}
}
func (p *Pipeline) Project(projections ...bson.E) bson.D {
	project := bson.D{}

	if len(projections) > 0 {
		project = bson.D{{"$project", bson.A{projections}}}
	}

	return project
}
