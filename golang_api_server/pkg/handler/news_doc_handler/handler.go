package news_doc_handler

import (
	"net/http"

	es_data_access "api-server/pkg/data_access/elasticsearch_data_access"
	"api-server/pkg/middleware"
	"api-server/pkg/services/news_doc_service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	SearchFunc SearchFunc
}

func NewHandler(esDAO es_data_access.ESDAO) *Handler {

	newsDocESDAO := news_doc_service.NewNewsDocESDAO(esDAO)

	return &Handler{
		SearchFunc: NewSearchFunc(newsDocESDAO),
	}
}

func (h *Handler) GetSearch() gin.HandlerFunc {

	return func(c *gin.Context) {

		// request and params binding
		reuqest := SearchRequest{}
		if err := c.ShouldBindQuery(&reuqest); err != nil {
			c.JSON(
				http.StatusBadRequest,
				middleware.NewErrorResponse(err),
			)
		}

		// process request
		response, err := h.SearchFunc(
			c.Request.Context(),
			&reuqest,
		)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				middleware.NewErrorResponse(err),
			)
		}

		// response
		c.JSON(
			http.StatusOK,
			middleware.NewResponse(response),
		)
	}
}
