package newshdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tainvecs/garage/apisrv/pkg/api/middleware"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

// Handler of news docs service
type Handler struct {
	SearchFunc SearchFunc
}

// New func creates a news docs service handler
func New(esDAO *esdao.DataAccessObject) *Handler {

	svcESDAO := newssvc.NewESDAO(esDAO)

	return &Handler{
		SearchFunc: NewSearchFunc(svcESDAO),
	}
}

// GetSearch handles search request
func (h *Handler) GetSearch() gin.HandlerFunc {

	return func(c *gin.Context) {

		// request and params binding
		reuqest := SearchRequest{}
		if err := c.ShouldBindQuery(&reuqest); err != nil {
			c.JSON(
				http.StatusBadRequest,
				middleware.NewErrorResponse(err),
			)
			return
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
			return
		}

		// response
		c.JSON(
			http.StatusOK,
			middleware.NewResponse(response),
		)
	}
}
