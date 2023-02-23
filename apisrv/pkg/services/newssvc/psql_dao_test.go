package newssvc_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/data_access/sqldao"
	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testPsqlDAO newssvc.PsqlDAO

func initPsql() {

	if os.Getenv("TEST_REAL") != "true" {
		return
	}

	// check if env missing
	dsn := os.Getenv("PSQL_NEWS_DSN")
	if len(strings.TrimSpace(dsn)) == 0 {
		panic(errors.New("missing env PSQL_NEWS_DSN"))
	}

	// init client
	client, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(errors.New("failed to init client from DSN"))
	}

	// new dao
	dao := newssvc.NewPsqlDAO(client)
	testPsqlDAO = dao
}

func TestPsqlDAOReal(t *testing.T) {

	if os.Getenv("TEST_REAL") != "true" {
		t.Skip()
	}

	initPsql()

	t.Run("subtestPsqlDAOGetReal", subtestPsqlDAOGetReal)
}

func subtestPsqlDAOGetReal(t *testing.T) {

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
