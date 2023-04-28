package db

import (
	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	DatabaseName string
	Conn         *mongo.Client
}

func ConnectNoSqlDb() MongoDatabase {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	mongoUrl := p.GetString("database.mongo.url", "")
	dbName := p.GetString("database.name", "")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl + dbName))
	if err != nil {
		panic(err)
	}

	return MongoDatabase{
		DatabaseName: "motion-go",
		Conn:         client,
	}
}
