package routes

import (
	"github.com/gin-gonic/gin"
	"todo/internal/api/http/v1/handlers"
)

func Setup(router *gin.Engine, h *handlers.Handler) {
	v1 := router.Group("/v1")
	r := v1.Group("/todo")

	// /v1/todo/
	r.GET("/", h.Get)
	r.POST("/", h.Create)
	r.PATCH("/:id", h.UpdateStatus)
	r.DELETE("/:id", h.Delete)
}
