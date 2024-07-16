package handler

import (
	new_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/new/usecases"
	user_usecase "github.com/ngoctb13/clean-architect-golang/internal/domains/user/usecases"

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
