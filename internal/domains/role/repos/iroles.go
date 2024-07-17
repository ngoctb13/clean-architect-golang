package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
)

type IRolesRepo interface {
	Create(ctx context.Context, record *models.Role) (*models.Role, error)
	GetByID(ctx context.Context, id int) (*models.Role, error)
	GetByName(ctx context.Context, name string) (*models.Role, error)
}
