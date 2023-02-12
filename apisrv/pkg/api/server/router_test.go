package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/handler"
	"github.com/tainvecs/garage/apisrv/pkg/api/server"
)

func TestSetUpRoute(t *testing.T) {

	// new mock handler
	mockHandler := handler.MockNewHandler()

	// init test router with mocked handler
	router := server.NewRouter()
	server.SetUpRoute(router, mockHandler)

	// test
	w := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/news-docs/v1/search?query=test", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
