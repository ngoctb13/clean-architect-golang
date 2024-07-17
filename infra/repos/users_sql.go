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
	err := u.db.Preload("Roles").First(&res, id).Error
	return res, err
}

func (u *usersSQLRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Preload("Roles").Where("username = ?", username).First(&user).Error
	return user, err
}

func (u *usersSQLRepo) AssignRole(ctx context.Context, userID int, roleID int) error {
	userRole := &models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return u.db.Create(userRole).Error
}
