package util

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabaseTest() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/test.db", getCurrentDirectory())), &gorm.Config{})
	if err != nil {
		panic("erro connection Db")
	}
	sqlDB, err := db.DB()

	return db, sqlDB
}

func RemoveDatabase() {
	err := os.Remove(fmt.Sprintf("%s/test.db", getCurrentDirectory()))
	if err != nil {
		panic(err)
	}
}
func getCurrentDirectory() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file info")
	}
	dir := path.Dir(filename)
	return dir
}
