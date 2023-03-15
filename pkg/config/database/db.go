package database

import (
	"database/sql"
	"fmt"
	"github.com/magiconair/properties"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, *sql.DB) {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	dbUsername := p.GetString("database.username", "")
	dbPassword := p.GetString("database.password", "")
	dbName := p.GetString("database.name", "")
	dbPort := p.GetInt("database.port", 0)
	dbHost := p.GetString("database.host", "")
	dsn := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s port=%d sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	//dsn := "host=localhost user=postgres password=frajolinha202 dbname=motion port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	if err != nil {
		panic("erro connection Db")
	}

	return db, sqlDB
}
