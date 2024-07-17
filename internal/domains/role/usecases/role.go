package usecases

import (
	"context"

	"github.com/ngoctb13/clean-architect-golang/internal/domain/models"
	"github.com/ngoctb13/clean-architect-golang/internal/domains/role/repos"
)

type Role struct {
	rolesRepo repos.IRolesRepo
}

func NewRole(rolesRepo repos.IRolesRepo) *Role {
	return &Role{
		rolesRepo: rolesRepo,
	}
}

func (r *Role) GetByID(ctx context.Context, id int) (*models.Role, error) {
	return r.rolesRepo.GetByID(ctx, id)
}

func (r *Role) Create(ctx context.Context, record *models.Role) (*models.Role, error) {
	return r.rolesRepo.Create(ctx, record)
}

func (r *Role) GetByName(ctx context.Context, name string) (*models.Role, error) {
	return r.rolesRepo.GetByName(ctx, name)
}
