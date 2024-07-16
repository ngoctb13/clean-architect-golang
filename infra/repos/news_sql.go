package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"

	"gorm.io/gorm"
)

type newsSQLRepo struct {
	db *gorm.DB
}

func InitNewsSQLRepo(db *gorm.DB) *newsSQLRepo {
	return &newsSQLRepo{
		db: db,
	}
}

func (n *newsSQLRepo) Create(ctx context.Context, record *models.New) (*models.New, error) {
	err := n.db.Create(record).Error
	return record, err
}

func (n *newsSQLRepo) Update(ctx context.Context, record *models.New) error {
	return n.db.Save(&models.New{ID: record.ID, Title: record.Title, Content: record.Content, AuthorID: record.AuthorID}).Error
}

func (n *newsSQLRepo) Delete(ctx context.Context, id int) error {
	return n.db.Delete(&models.New{}, id).Error
}

func (n *newsSQLRepo) GetByID(ctx context.Context, id int) (*models.New, error) {
	res := &models.New{}
	err := n.db.First(&res, id).Error
	return res, err
}

func (n *newsSQLRepo) GetByUserID(ctx context.Context, id int) ([]*models.New, error) {
	res := make([]*models.New, 0)
	err := n.db.Where("author_id = ?", id).Find(&res).Error
	return res, err
}
