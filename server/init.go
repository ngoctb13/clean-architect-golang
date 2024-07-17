package server

import (
	"github.com/gin-contrib/cors"
	hdl "github.com/ngoctb13/clean-architect-golang/handler"
	"github.com/ngoctb13/clean-architect-golang/infra/repos"
	new_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/new/usecases"
	role_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/role/usecases"
	user_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/user/usecases"
	"github.com/ngoctb13/clean-architect-golang/middleware"
)

type domains struct {
	user *user_usecase.User
	new  *new_usecase.New
	role *role_usecase.Role
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
	user := user_usecase.NewUser(repo.Users(), repo.Roles())
	new := new_usecase.InitNew(repo.News())
	role := role_usecase.NewRole(repo.Roles())

	return &domains{
		user: user,
		new:  new,
		role: role,
	}
}

func (s *Server) initRouter(domains *domains, repo repos.IRepo) {
	//init handler
	handler := hdl.NewHandler(domains.user, domains.new, domains.role)
	//api
	routerNoAuth := s.router.Group("v1")
	handlerNoAuth := handler
	handlerNoAuth.ConfigNoAuthRouteAPI(routerNoAuth)

	routerAuth := s.router.Group("v1")
	routerAuth.Use(middleware.AuthMiddleware())
	handler.ConfigAuthRouteAPI(routerAuth)
}
