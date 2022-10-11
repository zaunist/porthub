package echo

import (
	"github.com/gin-gonic/gin"
	"github.com/zaunist/porthub/pkg/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/echo", h.Get)
}

func (h *Handler) Get(c *gin.Context) {
	c.JSON(200, "hello")
}
