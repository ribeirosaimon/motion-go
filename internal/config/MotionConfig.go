package config

import (
	"context"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type MotionConfig struct {
	CacheTime    uint8
	ScrapingTime uint8
	HaveScraping bool
}

var config *MotionConfig

func NewMotionConfig(ctx context.Context, p *properties.Properties) *MotionConfig {
	if config == nil {
		config = &MotionConfig{}
		config.getConfigurations(ctx, p)
	}
	return config
}

func (m *MotionConfig) getConfigurations(ctx context.Context, p *properties.Properties) {
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
		config = &MotionConfig{
			CacheTime:    5,
			ScrapingTime: 3,
			HaveScraping: true,
		}
		collection.InsertOne(ctx, config)
	}
}

func GetConfiguration() MotionConfig {
	return *config
}
