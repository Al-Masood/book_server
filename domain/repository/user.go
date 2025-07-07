package repository

import (
	"context"
	"github.com/al-masood/book_server/domain/entity"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) error
	AuthenticateBasic(ctx context.Context, username string, password string) error
}
