package repository

import (
	"golang-gin-framework/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(DB *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		DB: DB,
	}
}

func (repository *BookRepositoryImpl) GetAll() ([]domain.Book, error) {
	var books []domain.Book

	err := repository.DB.Debug().Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (repository *BookRepositoryImpl) Create(book domain.Book) (domain.Book, error) {
	err := repository.DB.Debug().Create(&book).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) GetOne(id int) (domain.Book, error) {
	book := domain.Book{}

	err := repository.DB.Debug().Find(&book, id).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) Update(book domain.Book) (domain.Book, error) {
	err := repository.DB.Debug().Save(&book).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}

func (repository *BookRepositoryImpl) Delete(book domain.Book) (domain.Book, error) {
	err := repository.DB.Debug().Delete(&book).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book, nil
}
