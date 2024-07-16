package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ngoctb13/clean-architect-golang/handler/models"
	modelsUsecase "github.com/ngoctb13/clean-architect-golang/internal/domain/models"
)

const (
	UPDATE_SUCCESS_MESSAGE = "update success!"
	DELETE_SUCCESS_MESSAGE = "delete success!"
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

func (h *Handler) updateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &models.UpdateUserRequest{}
		if err := c.ShouldBind(req); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.user.Update(c, &modelsUsecase.User{
			ID:   req.ID,
			Name: req.Name,
			Age:  req.Age,
		})
		if err != nil {
			log.Printf("UpdateUser got error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": UPDATE_SUCCESS_MESSAGE})
	}
}

func (h *Handler) getUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		userID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.user.GetByID(c, userID)
		if err != nil {
			log.Printf("GetUserByID got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
