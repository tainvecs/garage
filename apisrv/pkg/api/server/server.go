package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// New func creates a new http server
func New(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

// Start the http server
// reference: https://github.com/gin-gonic/gin#manually
func Start(srv *http.Server, quit chan os.Signal) error {

	// initializing the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// gracefully shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
		return err
	}
	log.Println("Server exiting")

	return nil
}
