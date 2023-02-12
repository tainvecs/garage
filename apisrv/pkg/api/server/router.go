package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tainvecs/garage/apisrv/pkg/api/handler"
)

// NewRouter create a new router with gin Default
func NewRouter() *gin.Engine {
	return gin.Default()
}

// SetUpRoute func set up api routes with the handler
func SetUpRoute(
	router *gin.Engine,
	handler *handler.Handler,
) {

	// news doc
	newsDocGroup := router.Group("/news-docs/v1")
	{
		newsDocGroup.GET("/search", handler.NewsDocHandler.GetSearch())
	}
}
