package qmgo

import "go.mongodb.org/mongo-driver/mongo"

func (c *Client) Raw() *mongo.Client {
	return c.client
}
