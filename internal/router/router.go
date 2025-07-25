package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sancheschris/shorty/internal/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/shorten", controller.ShortenURL)
	}

	return r
}
