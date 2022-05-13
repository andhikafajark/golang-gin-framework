package service

import (
	"golang-gin-framework/dto/request"
	"golang-gin-framework/dto/response"
)

type BookService interface {
	GetAll() ([]response.BookResponse, error)
	Create(request request.CreateBook) (response.BookResponse, error)
	GetOne(id int) (response.BookResponse, error)
	Update(request request.UpdateBook) (response.BookResponse, error)
	Delete(id int) (response.BookResponse, error)
}
