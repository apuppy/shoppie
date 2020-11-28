package model

import (
	"fmt"
	"log"
	"shoppie/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
)

// mysql connection info
var (
	USER   = config.Conf.Database.User
	PASS   = config.Conf.Database.Password
	HOST   = config.Conf.Database.Host
	PORT   = config.Conf.Database.Port
	DBNAME = config.Conf.Database.DB
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
