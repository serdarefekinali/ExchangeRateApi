package database

import (
	"er-rabbit-consumer/model"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func DBConnect() (string, string) {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", host, user, dbname, password, dbport)
	return dialect, dbURI
}

func Connect(myStruct model.ExchangeRate) {
	dialect, _ := DBConnect()
	_, dbURI := DBConnect()
	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		panic(err.Error())
	}
	database := db.DB()
	err = database.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection to PostgreSQL was successful!")
	}
	defer db.Close()
	DB = db
	db.AutoMigrate(myStruct)
	db.Create(myStruct)
}
