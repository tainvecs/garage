package main

import (
	"os"

	"github.com/tainvecs/garage/apisrv/pkg/api/handler"
	"github.com/tainvecs/garage/apisrv/pkg/api/server"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
)

func main() {

	// read env
	esURL := os.Getenv("ES_URL")
	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")

	// init dao
	esDAO, err := esdao.New(esURL, esIndexIndex, esSearchIndex)
	if err != nil {
		panic(err)
	}

	// set-up router
	apiRouter := server.NewRouter()

	apiHandler := handler.NewHandler(esDAO)
	server.SetUpRoute(apiRouter, apiHandler)

	// start server
	quit := make(chan os.Signal, 1)
	apiServer := server.New(":80", apiRouter)
	err = server.Start(apiServer, quit)
	if err != nil {
		panic(err)
	}
}
