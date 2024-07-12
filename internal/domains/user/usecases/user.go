package usecases

import (
	"clean-arch-repo/internal/domain/models"
	"clean-arch-repo/internal/domains/user/repos"
	"context"
)

type User struct {
	usersRepo repos.IUsersRepo
}

func NewUser(usersRepo repos.IUsersRepo) *User {
	return &User{
		usersRepo: usersRepo,
	}
}

func (u *User) GetByID(ctx context.Context, id int) (*models.User, error) {
	return u.usersRepo.GetByID(ctx, id)
}

func (u *User) Create(ctx context.Context, record *models.User) (*models.User, error) {
	return u.usersRepo.Create(ctx, record)
}

func (u *User) Update(ctx context.Context, record *models.User) error {
	return u.usersRepo.Update(ctx, record)
}

func (u *User) Delete(ctx context.Context, id int) error {
	return u.usersRepo.Delete(ctx, id)
}
