package httpserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"

	"net/http"
	"time"
)

type Server interface {
	Run(router *gin.Engine) error
	Shutdown(ctx context.Context) error
}

type server struct {
	Port       string
	httpServer *http.Server
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *server) Run(handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           net.JoinHostPort("", s.Port),
		Handler:        handler, // http.HandlerFunc(slowHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	return s.httpServer.ListenAndServe()
}

func NewServer(port string) Server {
	return &server{
		Port: port,
	}
}