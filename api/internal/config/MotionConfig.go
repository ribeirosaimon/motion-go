package config

import (
	"context"
	"time"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/api/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type MotionConfig struct {
	CacheTime    uint8
	ScrapingTime uint8
	HaveScraping bool
	InitialValue float32
}

var config *MotionConfig

func NewMotionConfig(p *properties.Properties) *MotionConfig {
	if config == nil {
		config = &MotionConfig{}
		config.getConfigurations(p)
	}
	return config
}

func GetMotionConfig() MotionConfig {
	return *config
}

func (m *MotionConfig) getConfigurations(p *properties.Properties) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
	ctx.Done()
}

func GetConfiguration() MotionConfig {
	return *config
}
