package qmgo

import "go.mongodb.org/mongo-driver/mongo"

func (c *Collection) Raw() *mongo.Collection {
	return c.collection
}
