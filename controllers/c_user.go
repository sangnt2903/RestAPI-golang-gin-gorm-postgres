package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pipeline_security/models"
)

func GetUser(c *gin.Context){
	userId := c.Param("id")
	var user models.User
	//get User by ID

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	} else{
		var role models.Role

		models.DB.Model(&user).Related(&role)
		user.Role = role
		c.JSON(http.StatusOK, gin.H{
			"data" : user,
		})
		return
	}
}
