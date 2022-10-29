package api

import (
	"github.com/gin-gonic/gin"

	"github.com/zaunist/porthub/pkg/handler"
	"github.com/zaunist/porthub/pkg/handler/echo"
	"github.com/zaunist/porthub/pkg/handler/ports"
	"github.com/zaunist/porthub/pkg/handler/timeout"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	factories := []handler.RegisterFactory{
		echo.NewHandler,
		ports.NewHandler,
		timeout.NewHandler,
	}

	for i := range factories {
		h, err := factories[i]()
		if err != nil {
			panic(err)
		}
		h.ApplyRoute(r)
	}

	return r
}
