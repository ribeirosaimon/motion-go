package db

import (
	"database/sql"
	"fmt"
	"github.com/magiconair/properties"
	"gorm.io/driver/postgres"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var Conn *Connections

type Connections struct {
	sqlStruct
	nOSQLConn *mongo.Client
}

type sqlStruct struct {
	conn  *gorm.DB
	close *sql.DB
}

func (c *Connections) GetPostgreSQL() *gorm.DB {
	return c.sqlStruct.conn
}

func (c *Connections) ClosePostgreSQL() *sql.DB {
	return c.sqlStruct.close
}

func (c *Connections) InitializeDatabases() {
	c.connectSQL()
}

func (c *Connections) connectSQL() *gorm.DB {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
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
		return dbInstance
	}
	return c.sqlStruct.conn
}
