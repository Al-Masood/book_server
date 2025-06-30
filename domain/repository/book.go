package repository

import (
	"context"

	"github.com/al-masood/book_server/domain/entity"
)

type BookRepository interface {
	Create(ctx context.Context, book *entity.Book) error
	GetByID(ctx context.Context, id string) (*entity.Book, error)
	GetAll(ctx context.Context) ([]*entity.Book, error)
	Update(ctx context.Context, id string, book *entity.Book) error
	Delete(ctx context.Context, id string) error
}
