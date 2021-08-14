package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"

	//"github.com/gin-gonic/gin"
	"github.com/seinyan/gorest/config"

	"net/http"
	"time"
)

type Server interface {
	Run(router *gin.Engine) error
	Shutdown(ctx context.Context) error
}

type server struct {
	Addr       string
	config     *config.Config
	httpServer *http.Server
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *server) Run(router *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           net.JoinHostPort("", "8000"),
		Handler:        router, // http.HandlerFunc(slowHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	return s.httpServer.ListenAndServe()
}

func NewServer(conf *config.Config) Server {
	return &server{
		config: conf,
	}
}