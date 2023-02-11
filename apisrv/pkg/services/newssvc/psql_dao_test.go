package newssvc_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPsqlDAO() (newssvc.PsqlDAO, error) {

	// check if env missing
	dsn := os.Getenv("PSQL_DSN")
	if len(strings.TrimSpace(dsn)) == 0 {
		return nil, errors.New("missing env PSQL_DSN")
	}

	// init client
	client, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, errors.New("failed to init client from DSN")
	}

	// new dao
	dao := newssvc.NewPsqlDAO(client)

	return dao, nil
}

func TestPsqlDAOReal(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	t.Run("subtestPsqlDAOGetAllReal", subtestPsqlDAOGetAllReal)
}

func subtestPsqlDAOGetAllReal(t *testing.T) {

	ctx := context.Background()

	// init dao
	psqlDAO, err := InitPsqlDAO()
	assert.NoError(t, err)

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
	docs, err := psqlDAO.GetAll(ctx, &config)
	assert.NoError(t, err)
	assert.NotEmpty(t, docs)

	// print debug
	mar, err := json.MarshalIndent(docs, "", "\t")
	assert.NoError(t, err)
	fmt.Println(string(mar))
}
