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
	newsDocGroup := router.Group("/v1/news-docs")
	{
		// get docs
		// curl -XGET 'http://0.0.0.0/v1/news-docs?page=1&limit=5'
		newsDocGroup.GET("", handler.NewsDocHandler.Get())
		newsDocGroup.GET("/", handler.NewsDocHandler.Get())

		// search docs
		// curl -XGET 'http://0.0.0.0/v1/news-docs/search?query=test'
		newsDocGroup.GET("/search", handler.NewsDocHandler.GetSearch())
	}
}
