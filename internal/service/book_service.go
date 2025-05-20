package service

import (
	"github.com/nursultan/go-project-nurs/internal/models"
	"github.com/nursultan/go-project-nurs/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAllBooks() []models.Book {
	return s.repo.GetAll()
}

func (s *BookService) GetBookByID(id int) (models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) CreateBook(book models.Book) models.Book {
	return s.repo.Create(book)
}

func (s *BookService) UpdateBook(id int, book models.Book) (models.Book, error) {
	return s.repo.Update(id, book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.Delete(id)
}
