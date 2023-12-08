package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
	"logur.dev/logur"
)

type Option func(s *Server)

func WithTemporalClients(c *temporal.Clients) Option {
	return func(s *Server) {
		s.temporal = c
	}
}
func WithRouter(c *chi.Mux) Option {
	return func(s *Server) {
		s.router = c
	}
}

func WithConfig(cfg *Config) Option {
	return func(s *Server) {
		s.cfg = cfg
	}
}
func WithLogger(l logur.Logger) Option {
	return func(s *Server) {
		s.logger = l
	}
}

func WithTaskQueue(tq string) Option {
	return func(s *Server) {
		s.taskQueue = tq
	}
}
