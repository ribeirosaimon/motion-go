package db

import (
	"database/sql"
	"fmt"

	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var Conn *Connections

type Connections struct {
	sqlStruct
	noSqlStruct
}

type sqlStruct struct {
	conn  *gorm.DB
	close *sql.DB
}

type noSqlStruct struct {
	conn         *mongo.Client
	DatabaseName string
}

func (c *Connections) GetPostgreSQL() *gorm.DB {
	return c.sqlStruct.conn
}

func (c *Connections) GetMongoTemplate() *mongo.Client {
	return c.noSqlStruct.conn
}

func (c *Connections) ClosePostgreSQL() *sql.DB {
	return c.sqlStruct.close
}

func (c *Connections) InitializeDatabases(conf string) {
	c.connectSQL(conf)
	c.connectNoSQL(conf)
}

func (c *Connections) connectSQL(conf string) {
	p := properties.MustLoadFile(conf, properties.UTF8)
	dbUsername := p.GetString("database.username", "")
	dbPassword := p.GetString("database.password", "")
	dbName := p.GetString("database.name", "")
	dbPort := p.GetInt("database.port", 0)
	dbHost := p.GetString("database.host", "")
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%d sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)

	if c.sqlStruct.conn == nil {
		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		c.sqlStruct.conn = dbInstance
		close, err := dbInstance.DB()
		c.sqlStruct.close = close
	}
}

func (c *Connections) connectNoSQL(conf string) {
	p := properties.MustLoadFile(conf, properties.UTF8)
	mongoUrl := p.GetString("database.mongo.url", "")
	dbName := p.GetString("database.name", "")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl + dbName))
	if err != nil {
		panic(err)
	}
	c.noSqlStruct.conn = client
	c.noSqlStruct.DatabaseName = dbName
}
