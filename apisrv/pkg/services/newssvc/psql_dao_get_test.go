package newssvc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"
)

func TestPsqlDAOGetReal(t *testing.T) {

	if testPsqlDAO == nil {
		t.Skip()
	}

	ctx := context.Background()

	// query config
	config := sqldao.QueryConfig{
		Fields: []string{"id", "uuid", "link", "title"},
		Offset: 10,
		Limit:  10,
		PreloadAssociations: []string{
			newssvc.PsqlNewsAuthorsAssociation,
		},
	}

	// run get all
	docs, err := testPsqlDAO.Get(ctx, &config)
	assert.NoError(t, err)
	assert.NotEmpty(t, docs)

	// // print debug
	// mar, err := json.MarshalIndent(docs, "", "\t")
	// assert.NoError(t, err)
	// fmt.Println(string(mar))
}
