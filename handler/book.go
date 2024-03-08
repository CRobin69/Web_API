package handler

import (
	"fmt"
	"net/http"

	"robin-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Christopher Robin",
		"bio":  "Love to code and learn new things. I am a full stack developer and I love to work with Go, Python, and JavaScript",
	})
}
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "belajar Golang",
	})
}
func BookHandler(c *gin.Context) {
	id := c.Param("id")       //menangkap id yang diinput pada url
	title := c.Param("title") //menangkap title yang diinput pada url
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}
func QueryHandler(c *gin.Context) {
	title := c.Query("title") //bentuknya localhost:8080/query?title=belajar. query = ?
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) //errornya ditampung dlu
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages, //penampilan error message
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
