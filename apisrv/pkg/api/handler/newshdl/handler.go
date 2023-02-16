package newshdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tainvecs/garage/apisrv/pkg/api/middleware"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
	"gorm.io/gorm"
)

// Handler of news docs service
type Handler struct {
	GetFunc    GetFunc
	SearchFunc SearchFunc
}

// New func creates a news docs service handler
func New(psqlClient *gorm.DB, esDAO *esdao.DataAccessObject) *Handler {

	svcPsqlDAO := newssvc.NewPsqlDAO(psqlClient)
	svcESDAO := newssvc.NewESDAO(esDAO)

	return &Handler{
		GetFunc:    NewGetFunc(svcPsqlDAO),
		SearchFunc: NewSearchFunc(svcESDAO),
	}
}

// Get handles get news docs request
func (h *Handler) Get() gin.HandlerFunc {

	return func(c *gin.Context) {

		// request and params binding
		reuqest := GetRequest{}
		if err := c.ShouldBindQuery(&reuqest); err != nil {
			c.JSON(
				http.StatusBadRequest,
				middleware.NewErrorResponse(err),
			)
			return
		}

		// process request
		response, err := h.GetFunc(
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
