package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabase struct {
	DatabaseName string
	Conn         *mongo.Client
}
