package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var database *gorm.DB //database

func init() {

	err := godotenv.Load() //Load .env file
	if err != nil {
		panic(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	// Connection stringi yaratılır
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	database = conn
	database.Debug().AutoMigrate(&Account{}) //Database migration
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return database
}
