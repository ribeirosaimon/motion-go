package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabaseTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("erro connection Db")
	}
	
	return db
}
