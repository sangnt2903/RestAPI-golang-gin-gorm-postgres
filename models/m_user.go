package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName 	string 	`json:"firstName" gorm:"size(255)"`
	MiddleName	string 	`json:"middleName" gorm:"size(255)"`
	LastName	string 	`json:"lastName" gorm:"size(255)"`
	Point		int		`json:"point" gorm:"default:0"`
	Email		string 	`json:"email" gorm:"size(255)"`
	Password	string	`json:"password" gorm:"size(255)"`
	Role		Role	`gorm:"foreignkey:RoleId`
	RoleId		uint
}
