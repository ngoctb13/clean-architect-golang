package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"clean-arch-repo/handler/models"
	modelsUsecase "clean-arch-repo/internal/domain/models"
)

func (h *Handler) createUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.CreateUserRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.user.Create(c, &modelsUsecase.User{
			Name: request.Name,
			Age:  request.Age,
		})
		if err != nil {
			log.Printf("CreateUser got error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, res)
	}
}
