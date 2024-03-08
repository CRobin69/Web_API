package main

import (
	"fmt"
	"log"
	"robin-api/book"
	"robin-api/handler" // Replace "your-package-path" with the correct package path

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "postgres://bsltwhco:nO1kOk_uzD_1O7a8fonOwDUX6zNwAtU-@rain.db.elephantsql.com/bsltwhco"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db.AutoMigrate(&book.Book{}) //migrate database
	// //CRUD
	// book := book.Book{}
	// book.Title = "Belajar Goland"
	// book.Price = 100000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "Belajar Golang dari dasar"

	// err = db.Create(&book).Error //create data

	// if err != nil {
	// 	log.Fatal("Error creating book")
	// }
	var books [] book.Book // menampilkan semua data di dalam slice
	err = db.Debug().Find(&books).Error //read data
	if err != nil {
		log.Fatal("Error reading book")
	}
	for _,b := range books{
		fmt.Println(b.Title)
		fmt.Println(b)
	}
	router := gin.Default()
	v1 := router.Group("/v1") //grouping url / versioning api
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/hehe", handler.QueryHandler)
	v1.GET("/books/:id/:title", handler.BookHandler) //:id adalah parameter yang diinput pada url, dinamis
	v1.POST("/books", handler.PostBookHandler)
	router.Run()
}
