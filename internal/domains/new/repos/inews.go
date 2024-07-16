package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
)

type INewsRepo interface {
	Create(ctx context.Context, record *models.New) (*models.New, error)
	Update(ctx context.Context, record *models.New) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*models.New, error)
	GetByUserID(ctx context.Context, id int) ([]*models.New, error)
}
