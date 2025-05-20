package repository

import (
	"errors"
	"github.com/nursultan/go-project-nurs/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() []models.Book {
	var books []models.Book
	r.db.Find(&books)
	return books
}

func (r *BookRepository) GetByID(id int) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Book{}, errors.New("book not found")
	}
	return book, result.Error
}

func (r *BookRepository) Create(book models.Book) models.Book {
	r.db.Create(&book)
	return book
}

func (r *BookRepository) Update(id int, updatedBook models.Book) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Book{}, errors.New("book not found")
	}
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	r.db.Save(&book)
	return book, nil
}

func (r *BookRepository) Delete(id int) error {
	var book models.Book
	result := r.db.Delete(&book, id)
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}
