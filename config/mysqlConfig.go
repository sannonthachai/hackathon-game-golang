package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB(dbConfig string) {
	var err error
	db, err = gorm.Open("mysql", dbConfig)
	if err != nil {
		fmt.Println("Error connecting to Identity MySQL db: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
