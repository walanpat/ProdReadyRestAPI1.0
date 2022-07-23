package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

//NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new DB Connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Printf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable \n", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Err opening connection")
		fmt.Println(err)
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		fmt.Println("Err receiving ping")

		return db, err
	}
	return nil, nil
}
