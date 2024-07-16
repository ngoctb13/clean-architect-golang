package repos

import (
	new_repos "github.com/ngoctb13/clean-architect-golang/internal/domains/new/repos"
	user_repos "github.com/ngoctb13/clean-architect-golang/internal/domains/user/repos"
)

type IRepo interface {
	Users() user_repos.IUsersRepo
	News() new_repos.INewsRepo
}
