package mongodb

import (
	"context"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/gateway/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type enpointsRepository struct {
	ctx context.Context
	db  *mongo.Client
}

func NewEnpointsRepository(ctx context.Context, db *mongo.Client) model.EnpointsRepository {
	return &enpointsRepository{
		ctx: ctx,
		db:  db,
	}
}

func (er *enpointsRepository) FindByServicePrefix(servicePrefix string) (model.Enpoints, error) {

	collection := er.db.Database("anaksekolah").Collection("users")

	var enpoint model.Enpoints

	filter := bson.M{
		"service_prefix": servicePrefix,
	}

	err := collection.FindOne(er.ctx, filter).Decode(&enpoint)

	if err != nil {
		return model.Enpoints{}, err
	}

	return enpoint, nil
}
