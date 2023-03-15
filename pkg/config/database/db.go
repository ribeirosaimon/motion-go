package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, *sql.DB) {
	dsn := "host=localhost user=postgres password=security dbname=motion port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		panic("erro connection Db")
	}

	return db, sqlDB
}
