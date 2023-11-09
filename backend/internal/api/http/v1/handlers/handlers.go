package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/internal/entity"
	"todo/internal/usecase"
)

type Handler struct {
	uc usecase.UseCase
}

func NewHandler(uc usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Get(ctx *gin.Context) {
	items, err := h.uc.Get()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, items)
}

func (h *Handler) Create(ctx *gin.Context) {
	var item entity.Todo

	err := ctx.BindJSON(&item)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	err = h.uc.Create(item.Title)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
