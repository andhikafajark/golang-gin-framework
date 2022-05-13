package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-framework/controller"
	"golang-gin-framework/domain"
	"golang-gin-framework/repository"
	"golang-gin-framework/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/tutorial_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	err = db.AutoMigrate(&domain.Book{})
	if err != nil {
		log.Fatal("Auto Migrate error")
	}

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/", indexHandler)
	}

	r.GET("/books", bookController.Index)
	r.POST("/books", bookController.Create)
	r.GET("/books/:id", bookController.Show)
	r.PUT("/books/:id", bookController.Update)
	r.DELETE("/books/:id", bookController.Delete)

	_ = r.Run()
}

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test": "Test1",
	})
}
