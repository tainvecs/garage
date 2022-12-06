package router

import (
	"api-server/pkg/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	return gin.Default()
}

func SetUpRoute(
	router *gin.Engine,
	handler *handler.Handler,
) error {

	// news doc
	newsDocGroup := router.Group("/new-docs/v1")
	{
		newsDocGroup.GET("/search", handler.NewsDocHandler.GetSearch())
		newsDocGroup.GET("/search/", handler.NewsDocHandler.GetSearch())
	}

	return nil
}
