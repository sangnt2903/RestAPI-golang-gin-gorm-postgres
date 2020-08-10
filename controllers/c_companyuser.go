package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pipeline_security/models"
)

func GetCompanyUserByCompanyId(c *gin.Context){
	companyId := c.Param("id")
	var companyUsers []models.CompanyUser

	if err := models.DB.Where("company_id = ?", companyId).Find(&companyUsers).Error; err != nil{

		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}

	var result []models.CompanyUser

	for _,companyUser := range companyUsers {
		var user models.User
		var company models.Company
		var role	models.Role

		models.DB.Model(&companyUser).Related(&user)
		models.DB.Model(&user).Related(&role)
		models.DB.Model(&companyUser).Related(&company)

		user.Role = role
		companyUser.User = user
		companyUser.Company = company

		result = append(result, companyUser)
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : result,
	})

	return
}

func GetCompanyUserByUserID(c *gin.Context){
	userId := c.Param("id")
	var companyUser models.CompanyUser

	if err := models.DB.Where("user_id = ?", userId).First(&companyUser).Error; err != nil{

		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})

		return
	}

	var user models.User
	var company models.Company
	var role	models.Role

	models.DB.Model(&companyUser).Related(&user)
	models.DB.Model(&user).Related(&role)
	models.DB.Model(&companyUser).Related(&company)

	user.Role = role
	companyUser.User = user
	companyUser.Company = company



	c.JSON(http.StatusOK, gin.H{
		"data" : companyUser,
	})

	return
}
