package newssvc_test

import (
	"errors"
	"os"
	"strings"

	"github.com/tainvecs/garage/apisrv/pkg/services/newssvc"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testPsqlDAO newssvc.PsqlDAO

func init() {

	if os.Getenv("TEST_REAL") != "true" {
		return
	}

	// check if env missing
	dsn := os.Getenv("PSQL_DSN")
	if len(strings.TrimSpace(dsn)) == 0 {
		panic(errors.New("missing env PSQL_DSN"))
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
