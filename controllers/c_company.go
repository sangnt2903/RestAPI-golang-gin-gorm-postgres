package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pipeline_security/models"
)

func GetCompanies(c *gin.Context){
	var companies []models.Company

	models.DB.Find(&companies)

	c.JSON(http.StatusOK, gin.H{
		"data" : companies,
	})

	return
}

func GetCompany(c *gin.Context){
	companyId := c.Param("id")
	var company models.Company

	if err := models.DB.Where("id = ?", companyId).First(&company).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data" : company,
	})

	return
}

