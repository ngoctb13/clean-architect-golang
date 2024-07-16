package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ngoctb13/clean-architect-golang/handler/models"
	modelsUsecase "github.com/ngoctb13/clean-architect-golang/internal/domain/models"

	"github.com/gin-gonic/gin"
)

const (
	UPDATE_NEW_SUCCESS_MESSAGE = "update success"
	DELETE_NEW_SUCCESS_MESSAGE = "delete success"
)

func (h *Handler) createNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.CreateNewRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.new.Create(c, &modelsUsecase.New{
			Title:    request.Title,
			Content:  request.Content,
			AuthorID: request.Author,
		})
		if err != nil {
			log.Printf("CreateNew got error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *Handler) getNewByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		newId, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.new.GetByID(c, newId)
		if err != nil {
			log.Printf("GetNewByID got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *Handler) getNewsByAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("user_id")
		userID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := h.new.GetByUserID(c, userID)
		if err != nil {
			log.Printf("GetNewsByAuthor got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func (h *Handler) updateNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.UpdateNewRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.new.Update(c, &modelsUsecase.New{
			ID:       request.ID,
			Title:    request.Title,
			Content:  request.Content,
			AuthorID: request.Author,
		})
		if err != nil {
			log.Printf("UpdateNew got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": UPDATE_NEW_SUCCESS_MESSAGE})
	}
}

func (h *Handler) deleteNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param("id")
		newID, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = h.new.Delete(c, newID)
		if err != nil {
			log.Printf("DeleteNew got error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": DELETE_NEW_SUCCESS_MESSAGE})
	}
}
