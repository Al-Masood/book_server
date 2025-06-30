package inmemory

import (
	"context"
	"errors"
	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
)

type bookRepo struct {
	bookStorage map[string]*entity.Book
}

func NewBookRepository() repository.BookRepository {
	return &bookRepo{
		bookStorage: make(map[string]*entity.Book),
	}
}

func (r *bookRepo) Create(ctx context.Context, book *entity.Book) error {
	if _, exists := r.bookStorage[book.UUID]; exists {
		return errors.New("book already exists")
	}
	r.bookStorage[book.UUID] = book
	return nil
}

func (r *bookRepo) GetByID(ctx context.Context, id string) (*entity.Book, error) {
	book, exists := r.bookStorage[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (r *bookRepo) GetAll(ctx context.Context) ([]*entity.Book, error) {
	books := make([]*entity.Book, 0, len(r.bookStorage))
	for _, book := range r.bookStorage {
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepo) Update(ctx context.Context, id string, book *entity.Book) error {
	if _, exists := r.bookStorage[id]; !exists {
		return errors.New("book not found")
	}
	r.bookStorage[id] = book
	return nil
}

func (r *bookRepo) Delete(ctx context.Context, id string) error {
	if _, exists := r.bookStorage[id]; !exists {
		return errors.New("book not found")
	}
	delete(r.bookStorage, id)
	return nil
}