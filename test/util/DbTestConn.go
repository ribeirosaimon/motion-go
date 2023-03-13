package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabaseTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("erro connection Db")
	}
	//sqlDB, err := db.DB()
	//sqlDB.Close()
	return db
}
