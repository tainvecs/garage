package handler

import (
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"gorm.io/gorm"
)

// Handler for api server
type Handler struct {
	NewsDocHandler *newshdl.Handler
}

// NewHandler create a new api server Handler
func NewHandler(psqlClient *gorm.DB, esDAO *esdao.DataAccessObject) *Handler {

	newsDocHandler := newshdl.New(psqlClient, esDAO)

	return &Handler{
		NewsDocHandler: newsDocHandler,
	}
}
