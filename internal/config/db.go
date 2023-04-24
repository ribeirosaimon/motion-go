package config

import (
	"database/sql"
	"fmt"

	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func createDbInstance(dsn string) *gorm.DB {
	if dbInstance == nil {
		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return dbInstance
	}
	return dbInstance
}

func ConnectSqlDb() (*gorm.DB, *sql.DB) {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	dbUsername := p.GetString("database.username", "")
	dbPassword := p.GetString("database.password", "")
	dbName := p.GetString("database.name", "")
	dbPort := p.GetInt("database.port", 0)
	dbHost := p.GetString("database.host", "")
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%d sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	db := createDbInstance(dsn)
	sqlDB, err := db.DB()
	if err != nil {
		panic("erro connection Db")
	}

	return db, sqlDB
}

func ConnectNoSqlDb() *mongo.Client {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	mongoUrl := p.GetString("database.mongo.url", "")
	dbName := p.GetString("database.name", "")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl + dbName))
	if err != nil {
		panic(err)
	}
	return client
}
