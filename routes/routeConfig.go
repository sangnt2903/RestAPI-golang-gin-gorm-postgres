package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pipeline_security/controllers"
)

func SetRoutesConfig() *gin.Engine{

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message" : "Welcome",
		})
	})

	//#begin Book
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.PostBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.PutBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	//# end Book

	//#begin Author
	/*
	r.GET("/authors", controllers.GetAuthors)
	r.GET("/authors", controllers.PostAuthor)
	r.GET("/authors/:id", controllers.GetAuthor)
	r.GET("/authors/:id", controllers.PutAuthor)
	r.GET("/authors/:id", controllers.DeleteAuthor)
	*/
	//#end Author

	return r
}