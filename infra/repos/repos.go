package repos

import (
	"github.com/ngoctb13/clean-architect-golang/config"
	new_repos "github.com/ngoctb13/clean-architect-golang/internal/domains/new/repos"
	user_repos "github.com/ngoctb13/clean-architect-golang/internal/domains/user/repos"

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
