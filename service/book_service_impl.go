package service

import (
	"golang-gin-framework/domain"
	"golang-gin-framework/dto/request"
	"golang-gin-framework/dto/response"
	"golang-gin-framework/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (service *BookServiceImpl) GetAll() ([]response.BookResponse, error) {
	books, err := service.BookRepository.GetAll()

	var bookResponses []response.BookResponse

	for _, book := range books {
		bookResponse := convertToBookResponse(book)

		bookResponses = append(bookResponses, bookResponse)
	}

	return bookResponses, err
}

func (service *BookServiceImpl) Create(request request.CreateBook) (response.BookResponse, error) {
	book := domain.Book{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Rating:      request.Rating,
	}

	book, err := service.BookRepository.Create(book)

	bookResponse := convertToBookResponse(book)

	return bookResponse, err
}

func (service *BookServiceImpl) GetOne(id int) (response.BookResponse, error) {
	book, err := service.BookRepository.GetOne(id)

	bookResponse := convertToBookResponse(book)

	return bookResponse, err
}

func (service *BookServiceImpl) Update(request request.UpdateBook) (response.BookResponse, error) {
	book, _ := service.BookRepository.GetOne(request.Id)

	book.Title = request.Title
	book.Description = request.Description
	book.Price = request.Price
	book.Rating = request.Rating

	book, err := service.BookRepository.Update(book)

	bookResponse := convertToBookResponse(book)

	return bookResponse, err
}

func (service *BookServiceImpl) Delete(id int) (response.BookResponse, error) {
	book, _ := service.BookRepository.GetOne(id)

	book, err := service.BookRepository.Delete(book)

	bookResponse := convertToBookResponse(book)

	return bookResponse, err
}

func convertToBookResponse(book domain.Book) response.BookResponse {
	return response.BookResponse{
		Id:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
	}
}
