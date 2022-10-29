package timeout

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zaunist/porthub/pkg/handler"
)

type Handler struct{}

func NewHandler() (handler.RouteRegister, error) {
	return &Handler{}, nil
}

func (h *Handler) ApplyRoute(r *gin.Engine) {
	r.GET("/timeout", h.Get)
}

func (h *Handler) Get(c *gin.Context) {
	// wait for 15 second to test gateway reseponse
	time.Sleep(time.Second * 15)

	c.JSON(200, "timeout")
}
