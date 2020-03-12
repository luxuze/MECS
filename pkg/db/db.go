package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	sqlClient *gorm.DB
	err       error
)

func init() {
	db, err := gorm.Open("sqlite3", "mecs.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Code{})
}

// DB gormMysql 数据库链接
func DB() *gorm.DB {
	if sqlClient == nil || sqlClient.DB().Ping() != nil {
		sqlClient, err = gorm.Open("sqlite3", "mecs.db")
		sqlClient.AutoMigrate(&Code{})
		sqlClient.LogMode(true)
		if err != nil {
			log.Fatalf("Connect DB Failed: %v", err)
		}
	}
	return sqlClient
}
