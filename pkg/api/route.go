package api

import (
	"github.com/gin-gonic/gin"

	"github.com/zaunist/porthub/pkg/handler"
	"github.com/zaunist/porthub/pkg/handler/echo"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()

	factories := []handler.RegisterFactory{
		echo.NewHandler,
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
