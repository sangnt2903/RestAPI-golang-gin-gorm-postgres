package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	RoleName 	string	`json:"roleName" gorm:"size(255)"`
}
