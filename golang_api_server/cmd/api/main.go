package main

import (
	"api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/handler"
	"api-server/pkg/server"

	"os"
)

func main() {

	// read env
	esURL := os.Getenv("ES_URL")
	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")

	// init dao
	esDAO, err := elasticsearch_data_access.NewESDAO(esURL, esIndexIndex, esSearchIndex)
	if err != nil {
		panic(err)
	}

	// set-up router
	apiRouter := server.NewRouter()

	apiHandler := handler.NewHandler(*esDAO)
	server.SetUpRoute(apiRouter, apiHandler)

	// start server
	apiServer := server.NewServer(apiRouter)
	server.StartServer(apiServer)
}
