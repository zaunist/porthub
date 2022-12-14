package ports

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zaunist/porthub/pkg/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/server_port", h.Get)
}

func (h *Handler) Get(c *gin.Context) {
	port := strings.Split(c.Request.Host, ":")
	c.JSON(200, port[len(port)-1])
}
