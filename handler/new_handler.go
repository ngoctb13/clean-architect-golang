package handler

import (
	"clean-arch-repo/handler/models"
	modelsUsecase "clean-arch-repo/internal/domain/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
