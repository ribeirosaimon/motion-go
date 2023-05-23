package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

func (c *Connections) GetPgsqTemplate() *gorm.DB {
	return c.sqlStruct.conn
}

func (c *Connections) GetMongoTemplate() *mongo.Client {
	return c.noSqlStruct.conn
}

func (c *Connections) ClosePostgreSQL() *sql.DB {
	return c.sqlStruct.close
}

func (c *Connections) InitializeDatabases(conf *properties.Properties) {
	c.connectSQL(conf)
	c.connectNoSQL(conf)
}

func (c *Connections) connectSQL(p *properties.Properties) {
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

func (c *Connections) connectNoSQL(p *properties.Properties) {
	mongoUrl := p.GetString("database.mongo.url", "")
	dbName := p.GetString("database.name", "")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl + dbName))
	if err != nil {
		panic(err)
	}
	c.noSqlStruct.conn = client
	c.noSqlStruct.DatabaseName = dbName
}

func (c *Connections) InitializeTestDatabases(dir string) {
	dbDir := fmt.Sprintf("%s/test.db", dir)
	err := os.Remove(dbDir)
	if err != nil {
		fmt.Errorf("no have db")
	}

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/test.db", dir)), &gorm.Config{})
	if err != nil {
		panic("erro connection Db")
	}
	sqlDB, err := db.DB()

	c.sqlStruct.conn = db
	c.sqlStruct.close = sqlDB

}
