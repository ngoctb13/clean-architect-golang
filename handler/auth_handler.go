package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngoctb13/clean-architect-golang/handler/models"
	modelsUsecase "github.com/ngoctb13/clean-architect-golang/internal/domain/models"
	"github.com/ngoctb13/clean-architect-golang/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	MsgRegisterSuccess = "register success"
)

func (h *Handler) register() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.RegisterRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("failed to hash password: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &modelsUsecase.User{
			Username: request.Username,
			Password: string(hashedPassword),
			Name:     request.Name,
			Age:      request.Age,
		}

		createdUser, err := h.user.Create(c, user)
		if err != nil {
			log.Printf("create user got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = h.user.AssignRole(c, createdUser.ID, "user")
		if err != nil {
			log.Printf("failed to assign default role: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": MsgRegisterSuccess})
	}
}

func (h *Handler) login() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.LoginRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := h.user.GetByUsername(c, request.Username)
		if err != nil {
			log.Printf("GetByUsername got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			log.Printf("CompareHashAndPassword got error: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := utils.GenerateJWT(user)
		if err != nil {
			log.Printf("GenerateJWT got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
