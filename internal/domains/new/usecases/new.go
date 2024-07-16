package usecases

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
	"github.com/ngoctb13/clean-architect-golang/internal/domains/new/repos"
)

type New struct {
	newsRepo repos.INewsRepo
}

func InitNew(newsRepo repos.INewsRepo) *New {
	return &New{
		newsRepo: newsRepo,
	}
}

func (n *New) GetByID(ctx context.Context, id int) (*models.New, error) {
	return n.newsRepo.GetByID(ctx, id)
}

func (n *New) GetByUserID(ctx context.Context, id int) ([]*models.New, error) {
	return n.newsRepo.GetByUserID(ctx, id)
}

func (n *New) Create(ctx context.Context, record *models.New) (*models.New, error) {
	return n.newsRepo.Create(ctx, record)
}

func (n *New) Update(ctx context.Context, record *models.New) error {
	return n.newsRepo.Update(ctx, record)
}

func (n *New) Delete(ctx context.Context, id int) error {
	return n.newsRepo.Delete(ctx, id)
}
