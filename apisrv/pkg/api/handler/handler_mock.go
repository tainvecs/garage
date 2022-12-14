package handler

import (
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
)

// MockNewHandler create a new mocked api server Handler
func MockNewHandler() *Handler {

	mockNewsDocHandler := newshdl.MockNew()

	return &Handler{
		NewsDocHandler: mockNewsDocHandler,
	}
}
