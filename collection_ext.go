package qmgo

import (
	"context"

	opts "github.com/webx-top/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Collection) Raw() *mongo.Collection {
	return c.collection
}

// Find find by condition filter，return QueryI
func (c *Collection) FindId(ctx context.Context, id interface{}, opts ...opts.FindOptions) QueryI {
	return &Query{
		ctx:        ctx,
		collection: c.collection,
		filter:     bson.M{"_id": id},
		opts:       opts,
		registry:   c.registry,
	}
}
