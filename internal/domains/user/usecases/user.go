package usecases

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
	role_repos "github.com/ngoctb13/clean-architect-golang/internal/domains/role/repos"
	"github.com/ngoctb13/clean-architect-golang/internal/domains/user/repos"
)

type User struct {
	usersRepo repos.IUsersRepo
	rolesRepo role_repos.IRolesRepo
}

func NewUser(usersRepo repos.IUsersRepo, rolesRepo role_repos.IRolesRepo) *User {
	return &User{
		usersRepo: usersRepo,
		rolesRepo: rolesRepo,
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

func (u *User) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return u.usersRepo.GetByUsername(ctx, username)
}

func (u *User) AssignRole(ctx context.Context, userID int, roleName string) error {
	role, err := u.rolesRepo.GetByName(ctx, roleName)
	if err != nil {
		return err
	}
	return u.usersRepo.AssignRole(ctx, userID, role.ID)
}
