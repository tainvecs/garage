package handler

import (
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

// Resources needed by an handler
type Resources struct {
	NewsPsqlDAO newssvc.PsqlDAO
	NewsESDAO   newssvc.ESDAO
}

// Handler for api server
type Handler struct {
	NewsDocHandler *newshdl.Handler
}

// NewHandler create a new api server Handler
func NewHandler(resources *Resources) *Handler {

	newsDocHandler := newshdl.New(resources.NewsPsqlDAO, resources.NewsESDAO)

	return &Handler{
		NewsDocHandler: newsDocHandler,
	}
}
