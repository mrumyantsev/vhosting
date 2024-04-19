package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mrumyantsev/video-hosting/internal/config"
)

type Server struct {
	config *config.Config
	server *http.Server
}

func New(cfg *config.Config, router http.Handler) *Server {
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ServerReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(cfg.ServerWriteTimeoutSeconds) * time.Second,
		MaxHeaderBytes: cfg.ServerMaxHeaderBytes,
	}

	return &Server{
		config: cfg,
		server: srv,
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("cannot start server: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("cannot shut down the server: %w", err)
	}

	return nil
}
