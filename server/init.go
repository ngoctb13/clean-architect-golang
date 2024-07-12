package server

import (
	hdl "clean-arch-repo/handler"
	"clean-arch-repo/infra/repos"
	new_usecase "clean-arch-repo/internal/domains/new/usecases"
	user_usecase "clean-arch-repo/internal/domains/user/usecases"

	"github.com/gin-contrib/cors"
)

type domains struct {
	user *user_usecase.User
	new  *new_usecase.New
}

func (s *Server) initCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) initDomains(repo repos.IRepo) *domains {
	user := user_usecase.NewUser(repo.Users())
	new := new_usecase.InitNew(repo.News())

	return &domains{
		user: user,
		new:  new,
	}
}

func (s *Server) initRouter(domains *domains, repo repos.IRepo) {
	//init handler
	handler := hdl.NewHandler(domains.user, domains.new)

	//api
	routerAuth := s.router.Group("v1")
	handler.ConfigRouteAPI(routerAuth)
}
