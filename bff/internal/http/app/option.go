package app

import (
	"github.com/temporalio/temporal-demo-infra/bff/internal/clients/temporal"
)

type Option func(h *Handlers)

func WithTemporalClients(c *temporal.Clients) Option {
	return func(h *Handlers) {
		h.temporal = c
	}
}
func WithGeneratedAppDirectory(dir string) Option {
	return func(h *Handlers) {
		h.generatedAppDir = dir
	}
}
func WithMountPath(path string) Option {
	return func(h *Handlers) {
		h.mountPath = path
	}
}
