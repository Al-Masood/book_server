package service

import (
	"context"
	"errors"
	
	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
)

type BookService interface {
	Create(ctx context.Context, book *entity.Book) error
	GetByID(ctx context.Context, id string) (*entity.Book, error)
	GetAll(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, id string, book *entity.Book) error
	Delete(ctx context.Context, id string) error
}

type bookService struct {
	bookRepo repository.BookRepository
}

func (s *bookService) Create(ctx context.Context, book *entity.Book) error {
	if book == nil || book.UUID == "" || book.Name == "" {
		return errors.New("book, UUID, and name are required")
	}

	return s.bookRepo.Create(ctx, book)
}

func (s *bookService) GetByID(ctx context.Context, id string) (*entity.Book, error) {
	if id == "" {
		return nil, errors.New("ID is required")
	}

	return s.bookRepo.GetByID(ctx, id)
}

func (s *bookService) GetAll(ctx context.Context) ([]*entity.Book, error) {
	return s.bookRepo.GetAll(ctx)
}

func (s *bookService) Update(ctx context.Context, id string, book *entity.Book) error {
	if id == "" || book == nil || book.UUID == "" {
		return errors.New("valid ID and book are required")
	}

	if id != book.UUID {
		return errors.New("ID mismatch between path and book UUID")
	}

	return s.bookRepo.Update(ctx, id, book)
}

func (s *bookService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("ID is required")
	}

	return s.bookRepo.Delete(ctx, id)
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}