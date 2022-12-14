package middleware_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/middleware"
	"github.com/tainvecs/garage/apisrv/pkg/utils/strutils"
)

func TestNewSearchResonse(t *testing.T) {

	// test new Response
	ansRespStr := `
        {
            "data": [],
            "error": null
        }
    `
	ansRespStr = strutils.TrimAllIndent(ansRespStr)

	testResp := middleware.NewResponse([]string{})
	mar, err := json.Marshal(testResp)
	assert.NoError(t, err)
	testRespStr := string(mar)
	testRespStr = strutils.TrimAllIndent(testRespStr)

	assert.Equal(t, ansRespStr, testRespStr)

	// test new error Response
	ansErrRespStr := `
        {
            "data": null,
            "error": {"message":"a test error"}
        }
    `
	ansErrRespStr = strutils.TrimAllIndent(ansErrRespStr)

	testErrResp := middleware.NewErrorResponse(errors.New("a test error"))
	mar, err = json.Marshal(testErrResp)
	assert.NoError(t, err)
	testErrRespStr := string(mar)
	testErrRespStr = strutils.TrimAllIndent(testErrRespStr)

	assert.Equal(t, ansErrRespStr, testErrRespStr)
}
