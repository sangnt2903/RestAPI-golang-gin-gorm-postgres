package models

type Book struct {
	ID 		int 	`json: "id" gorm:"primary_key"`
	Name 	string	`json: "name"`
	Author	int 	`json: "author"`
}
