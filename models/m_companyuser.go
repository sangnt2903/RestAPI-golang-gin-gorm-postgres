package models

import "github.com/jinzhu/gorm"

type CompanyUser struct {
	gorm.Model
	Company		Company	`gorm:"foreignkey:CompanyId"`
	CompanyId	uint	`json:"companyId"; gorm:"index"`
	User		User	`gorm:"foreignkey:CompanyId"`
	UserId		uint	`json:"userId"; gorm:"index"`
	IsOwner		bool	`json:"isOwner" gorm:"type:boolean"`
}

