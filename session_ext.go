package qmgo

import "go.mongodb.org/mongo-driver/mongo"

func (s *Session) Raw() mongo.Session {
	return s.session
}
