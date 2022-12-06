package news_doc_handler

import (
	"net/http"

	"api-server/pkg/middleware"
	"api-server/pkg/services/news_doc_service"

	"github.com/gin-gonic/gin"
)

type NewsDocHandler struct {
	SearchFunc SearchFunc
}

func NewNewsDocHandler(esDAO news_doc_service.ESDAO) *NewsDocHandler {
	return &NewsDocHandler{
		SearchFunc: NewSearchFunc(esDAO),
	}
}

func (h *NewsDocHandler) GetSearch() gin.HandlerFunc {

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
