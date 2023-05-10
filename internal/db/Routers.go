package db

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Connections struct {
	SQLConn   *gorm.DB
	SQLClose  *sql.DB
	NOSQLConn *mongo.Client
}
