package main

import (
	"os"

	"github.com/tainvecs/garage/apisrv/pkg/api/handler"
	"github.com/tainvecs/garage/apisrv/pkg/api/server"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// read env
	newsPsqlDSN := os.Getenv("PSQL_NEWS_DSN")
	esURL := os.Getenv("ES_URL")
	esIndexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	esSearchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")

	// init dao and client
	client, err := gorm.Open(
		postgres.Open(newsPsqlDSN),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	newsPsqlDAO := newssvc.NewPsqlDAO(client)

	dao, err := esdao.New(esURL, esIndexIndex, esSearchIndex)
	if err != nil {
		panic(err)
	}
	newsESDAO := newssvc.NewESDAO(dao)

	// set-up router
	apiRouter := server.NewRouter()

	apiHandlerRes := handler.Resources{
		NewsPsqlDAO: newsPsqlDAO,
		NewsESDAO:   newsESDAO,
	}
	apiHandler := handler.NewHandler(&apiHandlerRes)
	server.SetUpRoute(apiRouter, apiHandler)

	// start server
	quit := make(chan os.Signal, 1)
	apiServer := server.New(":80", apiRouter)
	err = server.Start(apiServer, quit)
	if err != nil {
		panic(err)
	}
}
