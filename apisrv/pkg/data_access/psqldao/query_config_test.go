package psqldao_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/psqldao"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestApplyQueryConfig(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	// init db client
	client, err := gorm.Open(
		sqlite.Open("/Users/chlin/Projects/garage/apisrv/data/test-employee.sqlite.db"),
		&gorm.Config{},
	)
	assert.NoError(t, err)

	// apply query config and run
	conf := psqldao.QueryConfig{
		Fields:              []string{"test_field"},
		Limit:               10,
		Offset:              0,
		PreloadAssociations: []string{"test_association"},
	}
	conf.Apply(client)
}
