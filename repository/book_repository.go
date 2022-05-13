package repository

import "golang-gin-framework/domain"

type BookRepository interface {
	GetAll() ([]domain.Book, error)
	Create(book domain.Book) (domain.Book, error)
	GetOne(id int) (domain.Book, error)
	Update(book domain.Book) (domain.Book, error)
	Delete(domain.Book) (domain.Book, error)
}
