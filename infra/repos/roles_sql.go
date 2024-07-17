package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
	"gorm.io/gorm"
)

type roleSQLRepo struct {
	db *gorm.DB
}

func NewRoleSQLRepo(db *gorm.DB) *roleSQLRepo {
	return &roleSQLRepo{
		db: db,
	}
}

func (r *roleSQLRepo) Create(ctx context.Context, record *models.Role) (*models.Role, error) {
	err := r.db.Create(record).Error
	return record, err
}

func (r *roleSQLRepo) GetByID(ctx context.Context, id int) (*models.Role, error) {
	res := &models.Role{}
	err := r.db.First(&res, id).Error
	return res, err
}

func (r *roleSQLRepo) GetByName(ctx context.Context, name string) (*models.Role, error) {
	res := &models.Role{}
	err := r.db.Where("name = ?", name).First(&res).Error
	return res, err
}
