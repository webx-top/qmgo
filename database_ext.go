package qmgo

import "go.mongodb.org/mongo-driver/mongo"

func (db *Database) Raw() *mongo.Database {
	return db.database
}
