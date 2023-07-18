package config

import (
	"context"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type motionConfig struct {
	CacheTime    uint8
	ScrapingTime uint8
	HaveScraping bool
}

var config *motionConfig

func NewMotionConfig(ctx context.Context, p *properties.Properties) *motionConfig {
	if config == nil {
		config = &motionConfig{}
		config.getConfigurations(ctx, p)
	}
	return config
}

func (m *motionConfig) getConfigurations(ctx context.Context, p *properties.Properties) {
	collection := db.Conn.GetMongoTemplate().
		Database(p.GetString("database.mongo.name", "motion")).
		Collection("MotionConfig")

	filter := bson.M{}
	documents, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		panic(err)
	}
	if documents >= 1 {
		conn := collection.FindOne(ctx, bson.M{})
		conn.Decode(config)
	} else {
		config = &motionConfig{
			CacheTime:    5,
			ScrapingTime: 5,
			HaveScraping: true,
		}
		collection.InsertOne(ctx, config)
	}

}
