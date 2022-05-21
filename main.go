package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	fmt.Println("Database connection success")

	// ======
	// SERVICE
	// ======

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// ======
	// REPOSITORY
	// ======
	// bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()
	// book, err := bookRepository.FindByID(3)
	// book := book.Book{
	// 	Title:       "$100 Startup",
	// 	Description: "Good Book",
	// 	Price:       95000,
	// 	Rating:      4,
	// }
	// bookRepository.Create(book)

	// for _, book := range books {
	// fmt.Println("Title :", book.Title)
	// }

	// TODO: migrate the schema

	// ======
	// CREATE DATA
	// ======

	// book := book.Book{}
	// book.Title = "Atomic habits"
	// book.Price = 120000
	// book.Rating = 4
	// book.Description = "Buku self development"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error creating book")
	// 	fmt.Println("===================")
	// }

	// ======
	// READ DATA
	// ======
	// var book []book.Book

	// err = db.Debug().Find(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error creating book")
	// 	fmt.Println("===================")
	// }

	// for _, book := range book {
	// 	fmt.Println("Title : ", book)
	// }

	// ======
	// UPDATE DATA
	// ======
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error finding book")
	// 	fmt.Println("===================")
	// }

	// book.Title = "Man Tiger (Revised edition)"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error updating book")
	// 	fmt.Println("===================")
	// }

	// ======
	// DELETE DATA
	// ======
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error finding book")
	// 	fmt.Println("===================")
	// }

	// err = db.Delete(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error deleting book")
	// 	fmt.Println("===================")
	// }

	// fmt.Println("Title : %v", book)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	// v1.POST("/books", bookHandler.PostBooksHandler)

	// v1.GET("/", rootHandler)
	// v1.GET("/hello", helloHandler)
	// v1.GET("/books/:id/:title", booksHandler)
	// v1.GET("/query", queryHandler)
	// v1.POST("/books", postBooksHandler)

	// v2 := router.Group("/v2")

	router.Run(":8888")
}
