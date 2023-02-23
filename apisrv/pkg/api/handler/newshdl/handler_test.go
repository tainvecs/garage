package newshdl_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/esdao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPsqlClient() (*gorm.DB, error) {

	// check if env missing
	dsn := os.Getenv("PSQL_NEWS_DSN")
	if len(strings.TrimSpace(dsn)) == 0 {
		return nil, errors.New("missing env PSQL_NEWS_DSN")
	}

	// init client
	client, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func initESDAO() (*esdao.DataAccessObject, error) {

	// check if env missing
	url := os.Getenv("ES_URL")
	if len(strings.TrimSpace(url)) == 0 {
		return nil, errors.New("missing env ES_URL")
	}

	indexIndex := os.Getenv("ES_NEWS_DOC_INDEX_ALIAS")
	if len(strings.TrimSpace(indexIndex)) == 0 {
		return nil, errors.New("missing env ES_NEWS_DOC_INDEX_ALIAS")
	}

	searchIndex := os.Getenv("ES_NEWS_DOC_SEARCH_ALIAS")
	if len(strings.TrimSpace(searchIndex)) == 0 {
		return nil, errors.New("missing env ES_NEWS_DOC_SEARCH_ALIAS")
	}

	// es dao
	dao, err := esdao.New(url, indexIndex, searchIndex)
	if err != nil {
		return nil, err
	}

	return dao, nil
}

func TestNew(t *testing.T) {

	// init
	psqlClient, err := initPsqlClient()
	assert.NoError(t, err)
	svcPsqlDAO := newssvc.NewPsqlDAO(psqlClient)
	esDAO, err := initESDAO()
	svcESDAO := newssvc.NewESDAO(esDAO)
	assert.NoError(t, err)
	newsHandler := newshdl.New(svcPsqlDAO, svcESDAO)

	// setup
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/test-get", newsHandler.Get())
	router.GET("/test-search", newsHandler.GetSearch())
	respRecorder := httptest.NewRecorder()

	// test request
	req, err := http.NewRequest(http.MethodGet, "/test-get?page=1&limit=10", nil)
	assert.NoError(t, err)
	router.ServeHTTP(respRecorder, req)
	router.Use(newsHandler.Get())
	assert.Equal(t, 200, respRecorder.Code)

	req, err = http.NewRequest(http.MethodGet, "/test-search?query=test&page=1&limit=10", nil)
	assert.NoError(t, err)
	router.ServeHTTP(respRecorder, req)
	router.Use(newsHandler.GetSearch())
	assert.Equal(t, 200, respRecorder.Code)
}
