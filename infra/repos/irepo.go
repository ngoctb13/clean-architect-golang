package repos

import (
	new_repos "clean-arch-repo/internal/domains/new/repos"
	user_repos "clean-arch-repo/internal/domains/user/repos"
)

type IRepo interface {
	Users() user_repos.IUsersRepo
	News() new_repos.INewsRepo
}
