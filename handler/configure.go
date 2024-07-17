package handler

import (
	new_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/new/usecases"
	role_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/role/usecases"
	user_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/user/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	user *user_usecase.User
	new  *new_usecase.New
	role *role_usecase.Role
}

func NewHandler(user *user_usecase.User, new *new_usecase.New, role *role_usecase.Role) *Handler {
	return &Handler{
		user: user,
		new:  new,
		role: role,
	}
}

func (h *Handler) ConfigAuthRouteAPI(router *gin.RouterGroup) {
	//users
	router.POST("/users", h.createUser())
	router.PUT("/users/update", h.updateUser())
	router.GET("/users/:id", h.getUserByID())
	//news
	router.POST("/news", h.createNew())
	router.GET("/news/:id", h.getNewByID())
	router.GET("/news/user/:user_id", h.getNewsByAuthor())
	router.PUT("/news/update", h.updateNew())
	router.DELETE("/news/:id", h.deleteNew())
}

func (h *Handler) ConfigNoAuthRouteAPI(router *gin.RouterGroup) {
	router.POST("/login", h.login())
	router.POST("/register", h.register())
}
