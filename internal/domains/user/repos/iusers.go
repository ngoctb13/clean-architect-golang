package repos

import (
	"clean-arch-repo/internal/domain/models"
	"context"
)

type IUsersRepo interface {
	Create(ctx context.Context, record *models.User) (*models.User, error)
	Update(ctx context.Context, record *models.User) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*models.User, error)
}
