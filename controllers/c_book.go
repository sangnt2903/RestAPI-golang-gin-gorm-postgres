package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pipeline_security/models"
	"strconv"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"books" : books,
	})
	return
}

func GetBook(c *gin.Context) {
	var book models.Book
	bookID := c.Param("id")

	if err := models.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : http.StatusNotFound,
			"message" : "Record not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"book" : book,
	})

	return
}

func PostBook (c *gin.Context){

	// validate input
	var bookInput models.Book

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	if err := models.DB.Create(&bookInput).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	} else{
		c.JSON(http.StatusCreated, gin.H{
			"status"  : "created successfully",
			"message" : bookInput,
		})

		return
	}

}

func PutBook(c *gin.Context){
	bookID,_ := strconv.Atoi(c.Param("id"))
	var bookInput models.Book
	var bookCurrent models.Book

	// get book if exist
	if err := models.DB.Where("id = ?", bookID).First(&bookCurrent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Record not found",
		})

		return
	}

	//validate
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Bad request",
		})

		return
	}

	// check id
	if bookInput.ID != bookCurrent.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "Bad request",
		})
		return
	}

	//update
	if err := models.DB.Model(&bookCurrent).Update(bookInput).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error" : err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Modify Successfully!",
		})
	}

}

func DeleteBook (c *gin.Context) {
	var book models.Book
	bookId := c.Param("id")

	if err := models.DB.Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delete denied"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"status" : http.StatusOK,
		"message": "Delete successfully !",
	})

	return
}
