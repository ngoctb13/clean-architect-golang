package repos

import (
	"clean-arch-repo/config"
	new_repos "clean-arch-repo/internal/domains/new/repos"
	user_repos "clean-arch-repo/internal/domains/user/repos"

	"gorm.io/gorm"
)

type Repo struct {
	db  *gorm.DB
	cfg *config.MySQLConfig
}

func NewSQLRepo(db *gorm.DB, cfg *config.MySQLConfig) *Repo {
	return &Repo{
		db:  db,
		cfg: cfg,
	}
}

func (r *Repo) Users() user_repos.IUsersRepo {
	return NewUsersSQLRepo(r.db)
}

func (r *Repo) News() new_repos.INewsRepo {
	return InitNewsSQLRepo(r.db)
}
