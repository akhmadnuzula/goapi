package database

import (
	// "database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
  "gorm.io/gorm"
)

// var DB *sql.DB
var DB *gorm.DB

func InitDB() {
	// env
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	var err error
	// dataSourceName := "root:@tcp(127.0.0.1:3306)/goapi?parseTime=true"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
	// DB, err = sql.Open("mysql", dataSourceName)
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// err = DB.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Connected to database")
}