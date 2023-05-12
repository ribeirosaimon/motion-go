package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Connections struct {
	SQL
	NOSQLConn *mongo.Client
}

type SQL struct {
	Conn  *gorm.DB
	Close *sql.DB
}
