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
	InitialValue float32
}

var config *MotionConfig

func NewMotionConfig(ctx context.Context, p *properties.Properties) *MotionConfig {
	if config == nil {
		config = &MotionConfig{}
		config.getConfigurations(ctx, p)
	}
	return config
}

func GetMotionConfig() MotionConfig {
	return *config
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
			InitialValue: 1000,
		}
		collection.InsertOne(ctx, config)
	}
}

func GetConfiguration() MotionConfig {
	return *config
}
