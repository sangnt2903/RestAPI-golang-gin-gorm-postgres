
package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDataBase() {

	username := "postgres"
	password := "S@ng29031998"
	dbName := "pipeline_security"
	dbHost := "localhost"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	DB = db
}
