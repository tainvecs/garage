package server_test

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tainvecs/garage/apisrv/pkg/api/server"
)

func TestStartServer(t *testing.T) {

	// init
	router := server.NewRouter()
	srv := server.New(
		":http",
		router,
	)

	// start server and quit
	quit := make(chan os.Signal, 1)

	quit <- syscall.SIGTERM
	err := server.Start(srv, quit)
	assert.NoError(t, err)
}
