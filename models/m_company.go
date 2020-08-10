package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	CompanyName		string	`json:"companyName" gorm:"size(255)"`
	Location		string	`json:"location" gorm:"size(255)"`
}
