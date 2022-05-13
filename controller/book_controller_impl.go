package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-gin-framework/dto/request"
	"golang-gin-framework/service"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(service service.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		BookService: service,
	}
}

func (controller *BookControllerImpl) Index(c *gin.Context) {
	books, err := controller.BookService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (controller *BookControllerImpl) Create(c *gin.Context) {
	var createRequest request.CreateBook

	err := c.ShouldBindJSON(&createRequest)
	if err != nil {
		var errorMessages []string

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition :%s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := controller.BookService.Create(createRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (controller *BookControllerImpl) Show(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	book, err := controller.BookService.GetOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (controller *BookControllerImpl) Update(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	updateRequest := request.UpdateBook{
		Id: id,
	}

	err := c.ShouldBindJSON(&updateRequest)
	if err != nil {
		var errorMessages []string

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition :%s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := controller.BookService.Update(updateRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (controller *BookControllerImpl) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	book, err := controller.BookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
