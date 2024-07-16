package repos

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"

	"gorm.io/gorm"
)

type usersSQLRepo struct {
	db *gorm.DB
}

func NewUsersSQLRepo(db *gorm.DB) *usersSQLRepo {
	return &usersSQLRepo{
		db: db,
	}
}

func (u *usersSQLRepo) Create(ctx context.Context, record *models.User) (*models.User, error) {
	err := u.db.Create(record).Error
	return record, err
}

func (u *usersSQLRepo) Update(ctx context.Context, record *models.User) error {
	return u.db.Save(&models.User{ID: record.ID, Name: record.Name, Age: record.Age}).Error
}

func (u *usersSQLRepo) Delete(ctx context.Context, id int) error {
	return u.db.Delete(&models.User{}, id).Error
}

func (u *usersSQLRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	res := &models.User{}
	err := u.db.First(&res, id).Error
	return res, err
}
