package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"http_server/config"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	dbConf := config.GetDBConfig()
	var err error
	var dbName = dbConf.Get("database.path") + dbConf.Get("database.name") + ".db"
	//fmt.Println(dbName)
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
