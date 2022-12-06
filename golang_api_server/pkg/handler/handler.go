package handler

import (
	es_data_access "api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/handler/news_doc_handler"
)

type Handler struct {
	NewsDocHandler *news_doc_handler.Handler
}

func NewHandler(esDAO es_data_access.ESDAO) *Handler {

	newsDocHandler := news_doc_handler.NewHandler(esDAO)

	return &Handler{
		NewsDocHandler: newsDocHandler,
	}
}
