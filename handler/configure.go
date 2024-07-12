package handler

import (
	new_usecase "clean-arch-repo/internal/domains/new/usecases"
	user_usecase "clean-arch-repo/internal/domains/user/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	user *user_usecase.User
	new  *new_usecase.New
}

func NewHandler(user *user_usecase.User, new *new_usecase.New) *Handler {
	return &Handler{
		user: user,
		new:  new,
	}
}

func (h *Handler) ConfigRouteAPI(router *gin.RouterGroup) {
	//user
	router.POST("/users", h.createUser())

	//new
	router.POST("/news", h.createNew())
}
