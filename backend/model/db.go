package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// mysql connection info
const (
	USER   = "root"
	PASS   = "123456"
	HOST   = "127.0.0.1"
	PORT   = 3306
	DBNAME = "shoppie"
)

// Connect connect to mysql database
func Connect() *gorm.DB {
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", DSN)
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
