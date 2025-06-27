package inmemory

import (
	"context"
	"github.com/al-masood/book_server/domain/entity"
	"github.com/al-masood/book_server/domain/repository"
	"time"
)

type userRepository struct {
	// IN memory Map
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{
		// 		New map
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	/* implement me */
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	/* implement me */
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	/* implement me */

}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	/* implement me */
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	/* implement me */
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	/* implement me */
}
