package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbUri := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable TimeZone=Asia/Jakarta", username, password, dbName, dbPort, dbHost) //build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	// conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print("isi error : ", err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

//returns a handle tto the DB object
func GetDB() *gorm.DB {
	return db
}
