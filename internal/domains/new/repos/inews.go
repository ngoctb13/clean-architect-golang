package repos

import (
	"clean-arch-repo/internal/domain/models"
	"context"
)

type INewsRepo interface {
	Create(ctx context.Context, record *models.New) (*models.New, error)
	Update(ctx context.Context, record *models.New) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*models.New, error)
	GetByUserID(ctx context.Context, id int) ([]*models.New, error)
}
