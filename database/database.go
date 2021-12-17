package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"http_server/common"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func init() {
	dbConf := common.GetDBConfig()
	var err error
	var dbName = dbConf.Get("database.path") + dbConf.Get("database.name") + ".db"
	//fmt.Println(dbName)
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
