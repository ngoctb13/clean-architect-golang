package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
)

type IUsersRepo interface {
	Create(ctx context.Context, record *models.User) (*models.User, error)
	Update(ctx context.Context, record *models.User) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*models.User, error)
}
