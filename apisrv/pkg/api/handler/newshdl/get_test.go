package newshdl_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/handler/newshdl"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

func TestNewGetFunc(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	// prerequisite
	client, err := initPsqlClient()
	assert.NoError(t, err)
	svcPsqlDAO := newssvc.NewPsqlDAO(client)

	ctx := context.Background()
	getFunc := newshdl.NewGetFunc(svcPsqlDAO)

	// start testing
	request := newshdl.GetRequest{
		Page:  1,
		Limit: 5,
	}
	resp, err := getFunc(ctx, &request)
	assert.NoError(t, err)
	assert.Equal(t, resp.Total, 5)

	// // print debug
	// mar, err := json.MarshalIndent(resp, "", "\t")
	// assert.NoError(t, err)
	// fmt.Println(string(mar))
}
